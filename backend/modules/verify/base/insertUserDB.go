package base

import (
	"app/modules/tools"
	"log"
)

func InsertUserDB(body LoginBody) (id int) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		InsertUser,
		body.Email,
		body.Password,
	).Scan(
		&id,
	)

	if err != nil {
		log.Printf("failed to execute query InsertUser: %v", err)
	}

	return id
}

var InsertUser = `
insert into users (email, password)
	values ($1,$2)
	returning user_id
`
