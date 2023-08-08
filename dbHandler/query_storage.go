package dbHandler

/**
Storage of queries viable to execute into db
*/

const InsertNewGoal = `
	INSERT INTO goals
	(title, description, start_date, relevancy)
	VALUES
	(?, ?, ?, ?)
`
