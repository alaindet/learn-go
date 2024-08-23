package models

import (
	"app/core/db"
)

var createSql = `
	INSERT INTO "events" (
		"name",
		"description",
		"location",
		"date_time",
		"user_id"
	)
	VALUES
		(?, ?, ?, ?, ?)
`

func (e EventModel) Create() (EventModel, error) {

	stmt, err := db.DB.Prepare(createSql)
	if err != nil {
		return e, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return e, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return e, err
	}
	e.ID = id

	return e, nil
}
