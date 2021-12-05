package base

import (
	"app/modules/tools"
	"log"
)

func CheckMail(body SignupBody) (result bool) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		CHECKMailDB,
		&body.Email,
	).Scan(
		&result,
	)

	if err != nil {
		log.Printf("failed to execute query CHECKMailDB: %v", err)
	}

	return result
}

var CHECKMailDB = `
select exists (select email from users where email = $1)
`
