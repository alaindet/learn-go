package models

import (
	"app/core/db"
	"fmt"
)

var deleteSql = `DELETE FROM "events" WHERE "id" = ?`

func (e EventModel) Delete() (EventModel, error) {
	res, err := db.DB.Exec(deleteSql, e.ID)
	changedRows, err := res.RowsAffected()
	if err != nil || changedRows == 0 {
		return e, fmt.Errorf("Cannot delete event #%d", e.ID)
	}

	return e, nil
}
