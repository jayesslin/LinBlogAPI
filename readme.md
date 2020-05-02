## 构建linux可用app
```cassandraql 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
```
## linux 后台程序执行
```cassandraql
nohup ./appName &
```    

# LinBlogAPI现有接口
博客接口

     /api/blog/v1/publish/   //发布博客
     /api/blog/v1/blogs/  //获得博客列表
     /api/blog/v1/page  //单条博客查询
     /api/blog/v1/delete/   // 删除博客
     /api/blog/v1/categories/    //获得日志分类列表

## TODO
1. 更新博客//其实不需要
2. 获得博客分类列表 // 添加新表， 记录拆开 finish
3. 写追加内容
4. 根据分类筛选日志




