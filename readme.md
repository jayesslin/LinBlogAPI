## 构建linux可用app
```cassandraql 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
```
## linux 后台程序执行
```cassandraql
nohup ./appName &
```# LinBlogAPI
