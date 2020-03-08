package blog

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"golangDemo/dao"
	"golangDemo/model"
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
		log.Error("err!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Resulte": err,
		})
		return
	}
	log.Info("ctx.Request.body: %v", string(data))
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
