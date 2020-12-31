package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"warframes-tools-api/controller"
	"warframes-tools-api/model"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	model.SetupDb()

	r := gin.Default()

	warframe := r.Group("warframe")

	warframe.GET("", controller.GetWarframes)

	log.Fatal(r.Run(os.Getenv("PORT"))) // listen and serve on 0.0.0.0:8080 (
	// for windows "localhost:8080")
}
