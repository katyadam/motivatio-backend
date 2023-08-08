package api

import (
	"github.com/gin-gonic/gin"
	"motivatio-backend/datatypes"
	"motivatio-backend/dbHandler"
)

//for example i want to add new row into Goals table using post request from flutter

/**
Displaying GET/POST requests, that execute queries into db
*/

func displayRequests(r *gin.Engine) {
	db := dbHandler.GetDb()
	r.POST("/add-goal", func(c *gin.Context) {
		var goal datatypes.Goal
		if err := c.ShouldBindJSON(&goal); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//	Perform db insertion
		result, err := db.Exec(
			dbHandler.InsertNewGoal,
			goal.Title,
			goal.Description,
			goal.StartDate,
			goal.Relevancy)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		lastInsertID, _ := result.LastInsertId()
		goal.ID = int(lastInsertID)

		c.JSON(200, goal)
	})
}
