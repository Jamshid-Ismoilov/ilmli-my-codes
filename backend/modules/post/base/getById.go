package base

import (
	"app/modules/tools"
	"log"
)

func GetPostById(id int) (post FullPostBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		GETPostByIdDB,
		&id,
	).Scan(
		&post.ID,
		&post.Name,
		&post.Content,
		&post.Source,
		&post.Picture,
	)

	if err != nil {
		log.Printf("failed to execute query GETPostByIdDB: %v", err)
	}
	return post
}

var GETPostByIdDB = `
select
	post_id,
	name,
	content,
	source,
	picture
from posts
where post_id = $1

`
