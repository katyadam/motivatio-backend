package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"motivatio-backend/dbHandler"
)

// getting created database
var db *sql.DB = dbHandler.GetDb()

/*
*
Displaying GET/POST requests, that execute queries into db
*/
func displayRequests(r *gin.Engine) {
	r.POST("/add-goal", addGoal)
	r.GET("/goals", getGoals)
}
