package api

import (
	"fmt"
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
	err := r.Run(":8080")
	if err != nil {
		_ = fmt.Errorf("Error occurs: ", err)
	}
}
