package api

/**
Should store all GET/POST method handlers
*/

import (
	"github.com/gin-gonic/gin"
	"motivatio-backend/datatypes"
	"motivatio-backend/dbHandler"
	"time"
)

func getGoals(c *gin.Context) {
	rows, err := db.Query(dbHandler.SelectGoals)
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

		err := rows.Scan(&id, &title, &description, &startDate, &relevancy, &pinned, &done)
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
	result, err := db.Exec(
		dbHandler.InsertNewGoal,
		goal.Title,
		goal.Description,
		goal.StartDate,
		goal.Relevancy,
		goal.Pin,
		goal.Done)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	goal.ID = int(lastInsertID)

	c.JSON(200, goal)
}
