package dbHandler

/**
Storage of queries viable to execute into db
*/

const InsertNewGoal string = `
	INSERT INTO goals
	(title, description, start_date, relevancy, pin, done, user_id)
	VALUES
	($1, $2, $3, $4, $5, $6, $7)
`
const SelectGoals string = `
	SELECT * 
	FROM goals
	WHERE user_id = $1
`
const DeleteGoal string = `
	DELETE FROM goals WHERE id = $1
`
const InsertNewUser string = `
	INSERT INTO users
	(firstname, lastname, email, add_date, phone)
	VALUES
	($1, $2, $3, $4, $5)
`
const DeleteUser string = `
	DELETE FROM users WHERE id = $1
`
