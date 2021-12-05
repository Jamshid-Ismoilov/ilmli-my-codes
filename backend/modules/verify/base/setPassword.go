package base

import (
	"app/modules/tools"
	"log"
)

func SetPassword(body SignupBody, password string) (result string) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		SETPassword,
		&body.Email,
		&password,
	).Scan(
		&result,
	)

	if err != nil {
		log.Printf("failed to execute query SETPassword: %v", err)
	}
	return
}

var SETPassword = `
insert into passwords (email, password) values ($1, $2)
returning password_id
`
