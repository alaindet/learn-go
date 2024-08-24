package models

import (
	"app/core/db"
	"errors"

	"github.com/mattn/go-sqlite3"
)

var (
	ErrEventParticipationExists = errors.New("participation already exists")
)

var createParticipationSql = `
	INSERT INTO "event_participations" ("event_id", "user_id")
	VALUES (?, ?)
`

func (e EventModel) CreateParticipation(userId int64) (EventParticipation, error) {

	stmt, err := db.DB.Prepare(createParticipationSql)
	if err != nil {
		return EventParticipation{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.ID, userId)
	if err != nil {

		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
				return EventParticipation{}, ErrEventParticipationExists
			}
		}

		return EventParticipation{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return EventParticipation{}, err
	}

	var participation EventParticipation
	participation.ID = id
	participation.EventID = e.ID
	participation.UserID = userId

	return participation, nil
}
