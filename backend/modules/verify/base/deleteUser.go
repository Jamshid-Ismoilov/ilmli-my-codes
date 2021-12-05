package base

import (
	"app/modules/tools"
)

func DeleteUser(body LoginBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	_ = db.QueryRow(
		DeleteUserDB,
		body.Email,
		body.Password,
	)
}

var DeleteUserDB = `
delete from users where email = $1 and password = $2
`
