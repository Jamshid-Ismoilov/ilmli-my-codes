package base

import (
	"app/modules/tools"
	"log"
)

func GetByUserIdDB(id int) []PostBody {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	rows, err := db.Query(
		GETPostsByUserDB,
		&id,
	)

	if err != nil {
		log.Printf("failed to execute query GETPostsByUserDB: %v", err)
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

var GETPostsByUserDB = `
select
	p.post_id,
	p.name,
	p.picture
from users as u
join posts as p on p.user_id = u.user_id
where u.user_id = $1
`
