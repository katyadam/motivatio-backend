package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Display
/**
Displaying all viable GET/POST requests
*/
func Display() {
	r := gin.Default()
	r.Use(cors.Default())
	//	Start of api endpoints
	displayRequests(r)
	_ = r.Run(":8080")
}
