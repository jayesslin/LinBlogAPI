package blog

import (
	"blogapi/model"
	"blogapi/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func GetBlogs(ctx *gin.Context) {
	res, err := service.GetBlogService(ctx)
	if err != nil {
		log.Error("1get BLog failed")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Resulte": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
	return
}

func GetDetailBlog(ctx *gin.Context) {
	blogGet := model.Blog{}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Result": err,
		})
		return
	}
	err = json.Unmarshal(data, &blogGet)
	res, err := service.GetDetailBlogService(ctx, blogGet.BlogID)
	if err != nil {
		log.Error("Get One Blog failed")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Resulte": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
	return
}

func DeleteDetailBlog(ctx *gin.Context) {
	blogDelete := model.Blog{}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Result": err,
		})
		return
	}
	err = json.Unmarshal(data, &blogDelete)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Result": err,
		})
		return
	}
	err = service.DeleteDetailBlogService(ctx, blogDelete.BlogID)
	if err != nil {
		log.Error("Get One Blog failed")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Resulte": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete success",
	})
	return
}
func GetBlogsByTypes(ctx *gin.Context) {
	blogTypeGet := model.BlogType{}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Result": err,
		})
		return
	}
	err = json.Unmarshal(data, &blogTypeGet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Result": err,
		})
		return
	}
	res, err := service.GetBlogsByTypes(ctx, blogTypeGet.BlogType)
	if err != nil {
		log.Error("1get BLog failed")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Resulte": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
	return
}
