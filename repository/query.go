package repository

const (
	queryInsertUser = `
		INSERT INTO users
			(full_name, phone_number, password, salt)
		VALUES
			($1, $2, $3, $4)
		RETURNING 
			user_id, full_name, phone_number
	`

	queryGetUserPassword = `
		SELECT 
			user_id 
			, password
			, salt 
		FROM 
			users 
		WHERE 
			phone_number = $1
	`

	queryUpdateUserLoginCount = `
		UPDATE
			users
		SET 
			login_count = login_count + 1
		WHERE 
			phone_number = $1
	`

	queryGetUserByUserID = `
		SELECT 
			user_id
			, full_name 
			, phone_number
			, login_count
		FROM 
			users 
		WHERE 
			user_id = $1
	`
)
