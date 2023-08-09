package dbHandler

/**
Storage of queries viable to execute into db
*/

const InsertNewGoal string = `
	INSERT INTO goals
	(title, description, start_date, relevancy, pin, done)
	VALUES
	(?, ?, ?, ?, ?, ?)
`
const SelectGoals string = `
	SELECT * FROM goals
`
