package dbHandler

/**
Storage of queries viable to execute into db
*/

// GOALS
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

// USERS
const InsertNewUser string = `
	INSERT INTO users
	(firstname, lastname, email, add_date, phone)
	VALUES
	($1, $2, $3, $4, $5)
`
const DeleteUser string = `
	DELETE FROM users WHERE id = $1
`

// TAGS
const InsertNewTag string = `
	INSERT INTO tags
	(tag_name, color, user_id)
	VALUES
	($1, $2, $3)
`
const DeleteTag string = `
	DELETE FROM tags WHERE id = $1
`
const GetUserTags string = `
	SELECT * FROM tags where user_id = $1
`

// GOALS_TAGS
const AssignTagToGoal string = `
	INSERT INTO goals_tags
	(goal_id, tag_id)
	VALUES 
	($1, $2)
`
const GetGoalTags string = `
	SELECT * 
	FROM tags
	WHERE id in (
		SELECT tag_id 
		FROM goals_tags
		WHERE goal_id = $1
	)
`

// UPDATE QUERIES

// GOALS
const EditGoal string = `
	UPDATE goals
	SET title = $2,
	    description = $3
	WHERE id = $1
`
const ChangePinGoal string = `
	UPDATE goals
	SET pin = $2
	WHERE id = $1
`
const ChangeDoneGoal string = `
	UPDATE goals
	SET done = $2
	WHERE id = $1
`

// TAGS

const EditTag string = `
	UPDATE tags
	SET tag_name = $2,
	    color = $3
	WHERE id = $1
`
