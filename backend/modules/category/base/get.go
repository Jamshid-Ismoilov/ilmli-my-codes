package base

import (
	"app/modules/tools"
	"log"
)

func GetAllCategories() []CategoryBody {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	rows, err := db.Query(
		GETCategories,
	)

	defer rows.Close()

	if err != nil {
		log.Printf("facategoryiled to execute query GETCategories: %v", err)
	}

	var categories []CategoryBody

	for rows.Next() {
		var category CategoryBody

		rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.Picture,
		)

		categories = append(categories, category)
	}
	return categories
}

var GETCategories = `
select
	category_id,
	name,
	description,
	picture
from categories
`
