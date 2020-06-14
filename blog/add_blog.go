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

	res, err := service.PublicBlogService(ctx, blogAdded)

	if err != nil {
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
