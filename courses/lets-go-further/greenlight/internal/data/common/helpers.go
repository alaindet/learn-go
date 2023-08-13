package common

import (
	"context"
	"time"
)

func NewDatabaseContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}
