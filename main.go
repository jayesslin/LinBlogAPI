package main

import (
	"blogapi/db"
	"blogapi/logs"
	"blogapi/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("db err %v", err)
	}
	log.Info("hahahah")

	r := gin.Default()
	r.Use(util.Cors())
	r.Use(logs.LogerMiddleware())
	defer r.Run(":8000") // listen and serve on 0.0.0.0:8080
	register(r)
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
}
