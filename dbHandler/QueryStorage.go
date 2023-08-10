package dbHandler

/**
Storage of queries viable to execute into db
*/

const InsertNewGoal string = `
	INSERT INTO goals
	(title, description, start_date, relevancy, pin, done)
	VALUES
	($1, $2, $3, $4, $5, $6)
`
const SelectGoals string = `
	SELECT * FROM goals
`
const DeleteGoal string = `
	DELETE FROM goals WHERE id = $1
`
