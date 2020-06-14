package blog

import (
	"blogapi/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetBlogTypesCatagory(ctx *gin.Context) {
	res, err := service.GetBlogTypesCatagoryService(ctx)
	if err != nil {
		log.Error("1get BLog failed")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Result": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
	return
}
