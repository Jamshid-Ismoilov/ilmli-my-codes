package base

import (
	"app/modules/tools"
	"log"
)

func CheckSignupPassword(body LoginBody) (result bool) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		CHECKSignupPasswordDB,
		&body.Email,
		&body.Password,
	).Scan(
		&result,
	)

	if err != nil {
		log.Printf("failed to execute query CHECKSignupPasswordDB: %v", err)
	}

	return result
}

var CHECKSignupPasswordDB = `
select exists (select email from passwords where email = $1 and password = $2)
`
