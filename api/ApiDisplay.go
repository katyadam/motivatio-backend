package api

import (
	"github.com/gin-gonic/gin"
)

/*
Displaying GET/POST requests, that execute queries into database
*/
func displayRequests(r *gin.Engine) {
	r.POST("/add-goal", addGoal)
	r.GET("/goals/:userId", getGoals)
	r.POST("/delete-goal/:id", deleteGoal)

	r.POST("/add-user", addUser)
	r.POST("/delete-user/:id", deleteUser)

	r.POST("/add-tag", insertNewTag)
	r.POST("/delete-tag/:id", deleteTag)
	r.GET("/get-user-tags/:userId", getUserTags)
	r.GET("/get-goal-tags/:goalId", getGoalTags)
	r.POST("/assign-tag/:goalId/:tagId", assignTagToGoal)
}
