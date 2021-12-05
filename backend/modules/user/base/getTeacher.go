package base

import (
	"app/modules/tools"
	"log"
)

func GetTeacherById(id int) []TeacherBody {
	var DB tools.DB

	db := DB.ConnectionPQ()

	defer db.Close()

	rows, err := db.Query(
		GETTeachers,
		&id,
	)

	if err != nil {
		log.Printf("failed to execute query GETTeachers: %v", err)
	}

	defer rows.Close()
	var teachers []TeacherBody

	for rows.Next() {
		var teacher TeacherBody

		rows.Scan(
			&teacher.ID,
			&teacher.Fullname,
			&teacher.Highlight,
			&teacher.Picture,
		)

		teachers = append(teachers, teacher)
	}
	return teachers
}

var GETTeachers = `
select
	u.user_id,
	u.fullname,
	u.highlight,
	u.picture
from (
	select
		l.teacher_id as user_id
	from users as u
	join learning as l on l.user_id = u.user_id
	where u.user_id = $1
) as t
join users as u on u.user_id = t.user_id
`
