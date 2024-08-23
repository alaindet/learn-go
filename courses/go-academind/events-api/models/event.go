package models

import (
	"app/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e Event) Save() (Event, error) {
	sql := `
		INSERT INTO "events" (
			"name",
			"description",
			"location",
			"date_time",
			"user_id"
		)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(sql)
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

func GetAllEvents() ([]Event, error) {
	sql := `
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
	rows, err := db.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]Event, 0)

	for rows.Next() {
		var event Event
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

func GetEventByID(eventId string) (Event, error) {
	sql := `
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
	row := db.DB.QueryRow(sql, eventId)
	var event Event
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
