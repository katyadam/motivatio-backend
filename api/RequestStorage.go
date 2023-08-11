package api

/**
Should store all GET/POST method handlers
*/

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"motivatio-backend/datatypes"
	"motivatio-backend/dbHandler"
	"strconv"
	"time"
)

// importing database into package scope
var Db *sql.DB = dbHandler.GetDb()

func getGoals(c *gin.Context) {
	userId := c.Param("user_id")
	rows, err := Db.Query(dbHandler.SelectGoals, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer rows.Close() // po ukonceni teto funkce se provede tato funkce, lepsi try with resources

	var goals []map[string]interface{}

	for rows.Next() {
		var id int
		var title string
		var description string
		var startDate time.Time
		var relevancy float32
		var pinned bool
		var done bool
		var userId int

		err := rows.Scan(&id, &title, &description, &startDate, &relevancy, &pinned, &done, &userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		goal := map[string]interface{}{
			"id":          id,
			"title":       title,
			"description": description,
			"startDate":   startDate,
			"relevancy":   relevancy,
			"pin":         pinned,
			"done":        done,
			"userId":      userId,
		}
		goals = append(goals, goal)
	}
	c.JSON(200, goals)

}

func addGoal(c *gin.Context) {
	var goal datatypes.Goal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//	Perform db insertion
	result, err := Db.Exec(
		dbHandler.InsertNewGoal,
		goal.Title,
		goal.Description,
		goal.StartDate,
		goal.Relevancy,
		goal.Pin,
		goal.Done,
		goal.UserId)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	goal.ID = strconv.FormatInt(lastInsertID, 10)

	c.JSON(200, goal)
}

func deleteGoal(c *gin.Context) {
	goalId := c.Param("id")
	_, err := Db.Exec(dbHandler.DeleteGoal, goalId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	successMsg := "Goal with ID %s deleted successfully"
	c.JSON(200, gin.H{"status": fmt.Sprintf(successMsg, goalId)})
}

func addUser(c *gin.Context) {
	var user datatypes.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//	Perform db insertion
	result, err := Db.Exec(
		dbHandler.InsertNewUser,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.AddDate,
		user.Phone,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	user.ID = strconv.FormatInt(lastInsertID, 10)

	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	userId := c.Param("id")
	_, err := Db.Exec(dbHandler.DeleteUser, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	successMsg := "User with ID %s deleted successfully"
	c.JSON(200, gin.H{"status": fmt.Sprintf(successMsg, userId)})
}
