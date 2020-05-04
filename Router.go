package main

import (
	"blogapi/blog"
	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.Engine) {
	v1 := r.Group("/api/blog/v1")
	{
		v1.POST("/test/", blog.TestBlogV1)
		v1.POST("/publish/", blog.PublishBlog)
		v1.GET("/blogs/", blog.GetBlogs)
		v1.POST("/page", blog.GetDetailBlog)
		v1.POST("/deleteblog/", blog.DeleteDetailBlog)
		v1.GET("/categories/", blog.GetBlogTypesCatagory)
		v1.POST("/types/", blog.GetBlogsByTypes)
	}
}
