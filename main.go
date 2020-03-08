package main

import (
	"github.com/gin-gonic/gin"
	"golangDemo/db"
	"golangDemo/util"
	"log"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("db err %v", err)
	}
	r := gin.Default()
	r.Use(util.Cors())
	defer r.Run(":8001") // listen and serve on 0.0.0.0:8080
	register(r)
}
