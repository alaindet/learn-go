package models

import (
	"app/core/db"
)

var saveSql = `
	INSERT INTO "events" (
		"name",
		"description",
		"location",
		"date_time",
		"user_id"
	)
	VALUES (?, ?, ?, ?, ?)
`

func (e EventModel) Save() (EventModel, error) {

	stmt, err := db.DB.Prepare(saveSql)
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
