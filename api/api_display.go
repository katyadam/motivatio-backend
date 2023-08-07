package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Display() {
	r := gin.Default()
	r.Use(cors.Default())
	//	Start of api endpoints
	r.Run(":8080")
}
