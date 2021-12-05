package base

import (
	"app/modules/tools"
	"log"
)

func GetByCategoryIdDB(id int) []PostBody {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	rows, err := db.Query(
		GETPostsByCategoryDB,
		&id,
	)

	if err != nil {
		log.Printf("failed to execute query GETPostsByCategoryDB: %v", err)
	}

	defer rows.Close()
	var posts []PostBody

	for rows.Next() {
		var post PostBody

		rows.Scan(
			&post.ID,
			&post.Name,
			&post.Picture,
		)

		posts = append(posts, post)
	}
	return posts
}

var GETPostsByCategoryDB = `
select
	post_id,
	name,
	picture 
from posts
where category_id = $1
`
