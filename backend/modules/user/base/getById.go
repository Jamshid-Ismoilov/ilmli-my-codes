package base

import (
	"app/modules/tools"
	"log"
)

func GetUserById(id int) (user UserBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		GETWordById,
		&id,
	).Scan(
		&user.ID,
		&user.Fullname,
		&user.Highlight,
		&user.MobileNumber,
		&user.Bio,
		&user.Learners,
		&user.Picture,
	)

	if err != nil {
		log.Printf("failed to execute query GETWordById: %v", err)
	}
	return user
}

var GETWordById = `
select
	user_id,
	fullname,
	highlight,
	mobile_number,
	bio,
	learners,
	picture 
from users
where user_id = $1
`
