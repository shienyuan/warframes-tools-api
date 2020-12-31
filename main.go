package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"warframes-tools-api/controller"
	"warframes-tools-api/model"
)

func main() {
	model.SetupDb()

	r := gin.Default()

	warframe := r.Group("warframe")

	warframe.GET("", controller.GetWarframes)

	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
