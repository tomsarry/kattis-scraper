package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tomsarry/kattis-scraper/utils"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Static("/", "./public")
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://www.tomsarry.com", "http://localhost:3000"},
		AllowMethods: []string{"GET", "PUT", "POST"},
	}))

	r.POST("/kattis", utils.GetProblemsHandler)

	r.Run()
}
