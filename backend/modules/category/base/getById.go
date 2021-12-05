package base

import (
	"app/modules/tools"
	"log"
)

func GetCategoryById(id int) (category CategoryBody) {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	err := db.QueryRow(
		GETCategoryByIdDB,
		&id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.Picture,
	)

	if err != nil {
		log.Printf("facategoryiled to execute query GETCategoryByIdDB: %v", err)
	}
	return category
}

var GETCategoryByIdDB = `
select
	category_id,
	name,
	description,
	picture
from categories
where category_id = $1

`
