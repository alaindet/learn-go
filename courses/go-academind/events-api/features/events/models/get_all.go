package models

import (
	"app/core/db"
)

var getAllSql = `
	SELECT
		"id",
		"name",
		"description",
		"location",
		"date_time",
		"user_id"
	FROM
		"events"
`

func GetAll() ([]EventModel, error) {

	rows, err := db.DB.Query(getAllSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]EventModel, 0)

	for rows.Next() {
		var event EventModel
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, event)
	}

	return result, nil
}
