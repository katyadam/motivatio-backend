package api

import (
	"github.com/gin-gonic/gin"
)

/*
Displaying GET/POST requests, that execute queries into database
*/
func displayRequests(r *gin.Engine) {
	r.POST("/add-goal", addGoal)
	r.GET("/goals/:id", getGoals)
	r.POST("/delete-goal/:id", deleteGoal)

	r.POST("/add-user", addUser)
	r.POST("/delete-user/:id", deleteUser)

	r.POST("/add-tag", insertNewTag)
	r.POST("/delete-tag/:id", deleteTag)
	r.GET("/get-user-tags/:id", getUserTags)
	r.GET("/get-goal-tags/:id", getGoalTags)
	r.POST("/assign-tag/:goalId/:tagId", assignTagToGoal)

	r.POST("/edit-goal/:id/:title/:description", editGoal)
	r.POST("/change-goal-pin/:id/:pinState", changePinGoal)
	r.POST("/change-goal-done/:id/:isDone", changeDoneGoal)
	r.POST("edit-tag/:id/:tagName/:color", editTag)
}
