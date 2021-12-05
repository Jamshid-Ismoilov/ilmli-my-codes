package base

import (
	"app/modules/tools"
	"log"
)

func UpdatePost(body UpdatePostBody, id int) (post FullPostBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		UpdatePostDB,
		&body.ID,
		&body.Name,
		&body.Content,
		&body.Source,
		&body.Picture,
		&id,
	).Scan(
		&post.ID,
		&post.Name,
		&post.Content,
		&post.Source,
		&post.Picture,
	)

	if err != nil {
		log.Printf("failed to execute query UpdatePostDB: %v", err)
	}

	return
}

var UpdatePostDB = `
update posts
set name = $2, 
 	content = $3, 
  	source = $4, 
 	picture = $5 
where post_id = $1 and user_id = $6 
	returning post_id, name, content, source, picture
`
