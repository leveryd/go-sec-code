# Go漏洞靶场
"漏洞类型"包括：
* 任意文件读取
* 任意文件写入
* SSRF
* 命令执行注入
* SQL注入
* 模板注入
* 并发攻击

# 目录说明
* research: 比较少见的漏洞和研究性质的安全问题
* unsafe: 常见的安全问题
* safe: 对比"unsafe"安全的编码方式

# 资源
* [腾讯的Go安全编码指南](https://github.com/Tencent/secguide/blob/main/Go安全指南.md)
* [审计规则和工具](https://gist.github.com/leveryd/51b1ec0130d4b4e9df76d9413ae41239)
* [怎么做go安全研究](https://gist.github.com/leveryd/8581605b0f3532f8284bcfc4128f708c)

# 报告过的漏洞
* [CVE-2022-24863 swagger组件DoS](https://github.com/swaggo/http-swagger/security/advisories/GHSA-xg75-q3q5-cqmv)
* [CVE-2022-25757 json实现差异](https://www.openwall.com/lists/oss-security/2022/03/28/2)
* [CNVD-2022-51761 api/rpc框架gozero的dos问题](https://www.cnvd.org.cn/flaw/show/CNVD-2022-51761)

# 其他
// 不保证api接口稳定性，请自行测试

部分代码使用了 [copilot](https://github.com/github/copilot-docs) 自动生成。

