package base

import (
	"app/modules/tools"
	"log"
)

func UpdateUserDB(body UpdateUserBody, id int) (user UserBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		UpdatePostDB,
		&id,
		&body.Fullname,
		&body.Highlight,
		&body.MobileNumber,
		&body.Bio,
		&body.Picture,
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
		log.Printf("failed to execute query UpdatePostDB: %v", err)
	}

	return
}

var UpdatePostDB = `
update users
set fullname = $2, 
 	highlight = $3, 
  	mobile_number = $4, 
 	bio = $5, 
 	picture = $6 
where user_id = $1 
	returning user_id, fullname, highlight, mobile_number, bio, learners, picture
`
