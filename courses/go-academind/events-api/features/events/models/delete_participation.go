package models

import (
	"app/core/db"
	"errors"
	"fmt"
)

var (
	ErrEventParticipationMissing = errors.New("participation missing")
)

var getParticipationSql = `
	SELECT "id", "event_id", "user_id" FROM "event_participations"
	WHERE "event_id" = ? AND "user_id" = ?
`

var deleteParticipationSql = `
	DELETE FROM "event_participations"
	WHERE "id" = ?
`

func (e EventModel) DeleteParticipation(userId int64) (EventParticipation, error) {

	// Fetch existing participation
	row := db.DB.QueryRow(getParticipationSql, e.ID, userId)
	var participation EventParticipation
	err := row.Scan(
		&participation.ID,
		&participation.EventID,
		&participation.UserID,
	)
	if err != nil {
		return participation, ErrEventParticipationMissing
	}

	// Delete participation
	res, err := db.DB.Exec(deleteParticipationSql, participation.ID)
	changedRows, err := res.RowsAffected()
	if err != nil || changedRows == 0 {
		return participation, fmt.Errorf("Cannot delete participation to event #%d", e.ID)
	}

	return participation, nil
}
