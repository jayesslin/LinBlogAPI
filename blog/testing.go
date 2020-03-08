package blog

import (
	"github.com/gin-gonic/gin"
)

func TestBlogV1(c *gin.Context) {

	//
	//if err := db.Where("id = ?", id).First(&user).Error; err != nil {
	//	c.AbortWithStatus(http.StatusNotFound)
	//	fmt.Println(err)
	//} else {
	//	c.BindJSON(&user)
	//	db.Save(&user)
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": user,
	//	})
	//}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
