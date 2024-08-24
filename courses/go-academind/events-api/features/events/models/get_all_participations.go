package models

import (
	"app/core/db"
	usersModels "app/features/users/models"
)

var getAllParticipationsSql = `
	SELECT
		"u"."id",
		"u"."email"
	FROM
		"event_participations" AS "ep"
		INNER JOIN "users" AS "u" ON "ep"."user_id" = "u"."id"
	WHERE
		"ep"."event_id" = ?
`

func (e EventModel) GetParticipations() ([]usersModels.UserModelDisplay, error) {

	rows, err := db.DB.Query(getAllParticipationsSql, e.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]usersModels.UserModelDisplay, 0)

	for rows.Next() {
		var user usersModels.UserModelDisplay
		err := rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}
