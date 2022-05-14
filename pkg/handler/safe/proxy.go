package safe

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var noAvailableIP = errors.New("no available ip")
var noSupportedProtocol = errors.New("no supported protocol")

func isInternalIp(ip net.IP) bool {
	// 可能的更全的黑名单见 https://github.com/leveryd/go-sec-code/issues/3
	if ip.IsLoopback() || // 127.0.0.0/8
		ip.IsLinkLocalMulticast() || // 224.0.0.0/24
		ip.IsLinkLocalUnicast() { // 169.254.0.0/16
		return true
	}

	// 10.0.0.0/8
	// 172.16.0.0/12
	// 192.168.0.0/16
	if ip.IsPrivate() {
		return true
	}

	if ip4 := ip.To4(); ip4 != nil {
		return ip4[0] == 0 // 0.0.0.0/8
	}
	return false
}

func httpGet(xUrl string) ([]byte, error) {

	// 对于http请求，需要防止dns重绑定、防止30x跳转、检查是否请求内网资源
	// 自定义DialContext，使用"检查后的ip"建立tcp链接，防止dns重绑定
	// 自定义CheckRedirect，不允许30x跳转 （这里可能和实际业务需求不符合）
	// 需要保证请求头的Host与请求的url一致，而不是ip，否则可能会影响业务
	// https请求校验证书（这里可能和实际业务需求不符合）
	if strings.HasPrefix(xUrl, "http://") || strings.HasPrefix(xUrl, "https://") {
		customTransport := http.DefaultTransport.(*http.Transport).Clone()
		customTransport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			host, port, err := net.SplitHostPort(addr)
			// 解析host和 端口
			if err != nil {
				return nil, err
			}
			// dns解析域名
			ips, err := net.LookupIP(host)
			if err != nil {
				return nil, err
			}
			// 对所有的ip进行串行发起请求,选一个可用的
			for _, ip := range ips {
				fmt.Printf("%v -> %v is localip?: %v\n", addr, ip.String(), isInternalIp(ip))
				if isInternalIp(ip) {
					continue
				}
				// 拼接地址
				addr := net.JoinHostPort(ip.String(), port)
				con, err := net.Dial(network, addr)
				if err == nil {
					return con, nil
				}
			}

			return nil, noAvailableIP
		}
		client := http.Client{
			Transport: customTransport,
			CheckRedirect: func(req *http.Request, via []*http.Request) error { // 防止30x跳转
				return http.ErrUseLastResponse
			},
		}

		request, err := http.NewRequest("GET", xUrl, nil) // 添加host头
		if err != nil {
			return nil, err
		}

		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, noSupportedProtocol
}

func GoodHTTPGet(c *gin.Context) {
	xUrl := c.Query("url")
	get, err := httpGet(xUrl)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
	c.Writer.Write(get)
}

// [Go中的SSRF攻防战](https://segmentfault.com/a/1190000039009572)
// [golang中设置Host Header的小Tips](https://www.cnblogs.com/jinsdu/p/5161962.html)

// 需要兼顾业务功能和安全
