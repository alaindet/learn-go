package models

import (
	"app/core/db"
	"fmt"
)

var updateSql = `
	UPDATE "events"
	SET
		"name" = ?,
		"description" = ?,
		"location" = ?,
		"date_time" = ?
	WHERE
		"id" = ?
`

func (e EventModel) Update() (EventModel, error) {

	stmt, err := db.DB.Prepare(updateSql)
	if err != nil {
		return e, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.DateTime,
		e.ID,
	)

	changedRows, err := res.RowsAffected()
	if err != nil || changedRows == 0 {
		return e, fmt.Errorf("Cannot update event #%d", e.ID)
	}

	return e, nil
}
