package models

import (
	"time"
)

type EventModel struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

type EventParticipation struct {
	ID      int64
	EventID int64
	UserID  int64
}
