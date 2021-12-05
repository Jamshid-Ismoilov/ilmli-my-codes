package base

import (
	"app/modules/tools"
	"log"
)

func InsertPost(body NewPostBody, id int) (post FullPostBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		InsertPostDB,
		&body.Name,
		&body.Content,
		&body.Source,
		&body.Picture,
		&id,
		&body.CategoryID,
	).Scan(
		&post.ID,
		&post.Name,
		&post.Content,
		&post.Source,
		&post.Picture,
	)

	if err != nil {
		log.Printf("failed to execute query InsertUser: %v", err)
	}

	return
}

var InsertPostDB = `
insert into posts (name, content, source, picture, user_id, category_id)
	values ($1,$2,$3,$4,$5,$6)
	returning post_id, name, content, source, picture
`
