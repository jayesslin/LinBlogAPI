package main

import (
	"blogapi/db"
	"blogapi/util"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("db err %v", err)
	}

	r := gin.Default()
	r.Use(util.Cors())
	defer r.Run(":8000") // listen and serve on 0.0.0.0:8080
	register(r)
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
}
