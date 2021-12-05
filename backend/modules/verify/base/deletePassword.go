package base

import (
	"app/modules/tools"
)

func DeletePassword(body LoginBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	_ = db.QueryRow(
		DeletePasswordDB,
		body.Email,
	)
}

var DeletePasswordDB = `
delete from passwords where email = $1
`
