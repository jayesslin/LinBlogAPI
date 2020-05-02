## 构建linux可用app
```cassandraql 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
```
## linux 后台程序执行
```cassandraql
nohup ./appName &
```    

# LinBlogAPI现有接口
        博客
		v1.POST("/upload/", blog.PublishBlog) //发布博客
		v1.GET("/getblog/", blog.GetBlogs) //获得博客列表
		v1.POST("/page", blog.GetDetailBlog) //单条博客查询
		v1.POST("/deleteblog/",blog.DeleteDetailBlog) // 删除博客

## TODO
1. 更新博客
2. 获得博客分类列表 // 添加新表， 记录拆开
3. 写评论




