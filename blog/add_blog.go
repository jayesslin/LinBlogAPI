package blog

import (
	"blogapi/dao"
	"blogapi/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func PublishBlog(ctx *gin.Context) {
	blogAdded := model.AdditionBlog{}
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Resulte": err,
		})
		return
	}
	err = json.Unmarshal(data, &blogAdded)
	if err != nil {
		log.Error("err!", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Resulte": err,
		})
		return
	}
	log.Info("ctx.Request.body: %v", string(data))
	//添加博客类型
	go func() {
		err := dao.BlogTypedaoOnceInstance().AddBlogType(ctx, blogAdded)
		if err != nil {
			log.Error("Add BLog Type failed")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Result": err,
			})
			return
		}
	}()
	//添加博客记录
	res, err := dao.BlogdaoOnceInstance().AddBlog(ctx, blogAdded)
	if err != nil {
		log.Error("Add BLog failed")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Result": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Result": "success",
		"res":    res,
	})
	return
}
