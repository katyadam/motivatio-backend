package api

import (
	"github.com/gin-gonic/gin"
)

/*
Displaying GET/POST requests, that execute queries into database
*/
func displayRequests(r *gin.Engine) {
	r.POST("/add-goal", addGoal)
	r.GET("/goals", getGoals)
	r.POST("/delete-goal/:id", deleteGoal)
}
