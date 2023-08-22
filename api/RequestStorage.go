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
	userId := c.Param("id")
	rows, err := Db.Query(dbHandler.SelectGoals, userId)
	handleError(err, 500, c)
	defer rows.Close() // po ukonceni teto funkce se provede tato funkce, lepsi try with resources

	var goals []map[string]interface{}

	for rows.Next() {
		var id string
		var title string
		var description string
		var startDate time.Time
		var relevancy float32
		var pinned bool
		var done bool
		var userId string

		err := rows.Scan(&id, &title, &description, &startDate, &relevancy, &pinned, &done, &userId)
		handleError(err, 500, c)

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
	err := c.ShouldBindJSON(&goal)
	handleError(err, 400, c)
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

	handleError(err, 500, c)

	lastInsertID, _ := result.LastInsertId()
	goal.ID = strconv.FormatInt(lastInsertID, 10)

	c.JSON(200, goal)
}

func deleteGoal(c *gin.Context) {
	goalId := c.Param("id")
	_, err := Db.Exec(dbHandler.DeleteGoal, goalId)
	handleError(err, 500, c)

	successMsg := "Goal with ID %s deleted successfully"
	c.JSON(200, gin.H{"status": fmt.Sprintf(successMsg, goalId)})
}

func addUser(c *gin.Context) {
	var user datatypes.User
	err := c.ShouldBindJSON(&user)
	handleError(err, 400, c)
	//	Perform db insertion
	result, err := Db.Exec(
		dbHandler.InsertNewUser,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.AddDate,
		user.Phone,
	)
	handleError(err, 500, c)

	lastInsertID, _ := result.LastInsertId()
	user.ID = strconv.FormatInt(lastInsertID, 10)

	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	userId := c.Param("id")
	_, err := Db.Exec(dbHandler.DeleteUser, userId)
	handleError(err, 500, c)

	successMsg := "User with ID %s successfully delete!"
	c.JSON(200, gin.H{"status": fmt.Sprintf(successMsg, userId)})
}

func insertNewTag(c *gin.Context) {
	var tag datatypes.Tag
	err := c.ShouldBindJSON(&tag)
	handleError(err, 400, c)

	result, err := Db.Exec(
		dbHandler.InsertNewTag,
		tag.TagName,
		tag.Color,
		tag.UserId,
	)
	handleError(err, 500, c)

	lastInsertID, _ := result.LastInsertId()
	tag.ID = strconv.FormatInt(lastInsertID, 10)

	c.JSON(200, tag)
}

func deleteTag(c *gin.Context) {
	id := c.Param("id")
	_, err := Db.Exec(dbHandler.DeleteTag, id)
	handleError(err, 500, c)

	successMsg := "Tag with ID %s successfully deleted!"
	c.JSON(200, gin.H{"status": fmt.Sprintf(successMsg, id)})
}

func getUserTags(c *gin.Context) {
	userId := c.Param("id")
	rows, err := Db.Query(dbHandler.GetUserTags, userId)
	handleError(err, 500, c)

	var tags []map[string]interface{}

	for rows.Next() {
		var id int
		var tagName string
		var color int
		var userId int

		err := rows.Scan(&id, &tagName, &color, &userId)
		handleError(err, 500, c)

		tag := map[string]interface{}{
			"id":      id,
			"tagName": tagName,
			"color":   color,
			"userId":  userId,
		}
		tags = append(tags, tag)
	}
	c.JSON(200, tags)
}

func assignTagToGoal(c *gin.Context) {
	goalId, tagId := c.Param("goalId"), c.Param("tagId")
	_, err := Db.Exec(dbHandler.AssignTagToGoal, goalId, tagId)
	handleError(err, 500, c)

	successMsg := "TagId: %s assigned to GoalId: %s"
	c.JSON(200, gin.H{"status": fmt.Sprintf(successMsg, tagId, goalId)})
}

func getGoalTags(c *gin.Context) {
	goalId := c.Param("id")

	rows, err := Db.Query(dbHandler.GetGoalTags, goalId)
	handleError(err, 500, c)

	var tags []map[string]interface{}

	for rows.Next() {
		var id int
		var tagName string
		var color int
		var userId int

		err := rows.Scan(&id, &tagName, &color, &userId)
		handleError(err, 500, c)

		tag := map[string]interface{}{
			"id":      id,
			"tagName": tagName,
			"color":   color,
			"userId":  userId,
		}
		tags = append(tags, tag)
	}
	c.JSON(200, tags)
}

func editGoal(c *gin.Context) {
	goalId, title, description := c.Param("id"), c.Param("title"), c.Param("description")

	_, err := Db.Exec(dbHandler.EditGoal, goalId, title, description)
	handleError(err, 500, c)

	msg := fmt.Sprintf("GoalId: %s has been edited!", goalId)
	c.JSON(200, msg)
}

func changePinGoal(c *gin.Context) {
	goalId, pinState := c.Param("id"), c.Param("pinState")

	_, err := Db.Exec(dbHandler.ChangePinGoal, goalId, pinState)
	handleError(err, 500, c)

	msg := fmt.Sprintf("GoalId's: %s state of being pinned change!", goalId)
	c.JSON(200, msg)
}

func changeDoneGoal(c *gin.Context) {
	goalId, isDone := c.Param("id"), c.Param("isDone")

	_, err := Db.Exec(dbHandler.ChangeDoneGoal, goalId, isDone)
	handleError(err, 500, c)

	msg := fmt.Sprintf("GoalId's: %s state of being done changed!", goalId)
	c.JSON(200, msg)
}

func editTag(c *gin.Context) {
	tagId, tagName, color := c.Param("id"), c.Param("tagName"), c.Param("color")

	_, err := Db.Exec(dbHandler.EditTag, tagId, tagName, color)
	handleError(err, 500, c)

	msg := fmt.Sprintf("TagId: %s has been edited", tagId)
	c.JSON(200, msg)
}
