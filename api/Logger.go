package api

import "github.com/gin-gonic/gin"

func handleError(err error, code int, c *gin.Context) {
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
	}
}
