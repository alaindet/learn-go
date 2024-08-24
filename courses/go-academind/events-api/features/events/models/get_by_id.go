package models

import (
	"app/core/db"
)

var getByIdSql = `
	SELECT
		"id",
		"name",
		"description",
		"location",
		"date_time",
		"user_id"
	FROM
		"events"
	WHERE
		"id" = ?
`

func GetByID(eventId string) (EventModel, error) {

	row := db.DB.QueryRow(getByIdSql, eventId)

	var event EventModel
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	)

	if err != nil {
		return event, err
	}

	return event, nil
}
