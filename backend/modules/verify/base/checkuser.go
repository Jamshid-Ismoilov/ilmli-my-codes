package base

import (
	"app/modules/tools"
	"log"
)

func CheckUserDB(body LoginBody) (result int) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		CHECKUserDB,
		&body.Email,
		&body.Password,
	).Scan(
		&result,
	)

	if err != nil {
		log.Printf("failed to execute query CHECKUserDB: %v", err)
	}

	return result
}

var CHECKUserDB = `
select user_id from users where email = $1 and password = $2
`
