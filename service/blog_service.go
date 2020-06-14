package service

import (
	"blogapi/dao"
	"blogapi/model"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
)

func PublicBlogService(ctx *gin.Context, blogAdded model.AdditionBlog) (*model.AdditionBlog, error) {
	//添加博客类型
	err := dao.BlogTypedaoOnceInstance().AddBlogType(ctx, blogAdded)
	if err != nil {
		log.Error("Add BLog Type failed")
		return nil, err
	}
	//添加博客记录
	res, err := dao.BlogdaoOnceInstance().AddBlog(ctx, blogAdded)
	if err != nil {
		log.Error("Add BLog failed")
		return nil, err
	}

	return &res, nil
}
func GetBlogService(ctx *gin.Context) ([]*model.Blog, error) {
	res, err := dao.BlogdaoOnceInstance().GetBlogs(ctx, 1, 1)
	if err != nil {
		log.Error("1get BLog failed")
		return nil, err
	}
	return res, nil
}
func GetDetailBlogService(ctx *gin.Context, blogID int) (*model.Blog, error) {
	res, err := dao.BlogdaoOnceInstance().GetSpecifyBlog(ctx, blogID)
	if err != nil {
		log.Error("Get One Blog failed")
		return nil, err
	}
	return res, nil
}
func DeleteDetailBlogService(ctx *gin.Context, blogID int) error {
	err := dao.BlogdaoOnceInstance().DeletBlog(ctx, blogID)
	if err != nil {
		log.Error("Get One Blog failed")
		return err
	}
	return nil
}
func GetBlogsByTypes(ctx *gin.Context, blogType string) ([]*model.Blog, error) {
	if blogType == "" {
		blogType = "个人"
	}
	res, err := dao.BlogdaoOnceInstance().GetBlogsByType(ctx, 1, 1, blogType)
	if err != nil {
		log.Error("1get BLog failed")
		return nil, err
	}
	return res, nil
}
func GetBlogTypesCatagoryService(ctx *gin.Context) ([]*model.BlogTypeCategory, error) {
	res, err := dao.BlogTypedaoOnceInstance().GetBlogTypeCategory(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
