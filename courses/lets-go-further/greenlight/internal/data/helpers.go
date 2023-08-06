package data

import (
	"context"
	"time"
)

func NewDatabaseContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

// func NewDatabaseContextFrom(
// 	parentCtx *context.Context,
// 	timeout time.Duration,
// ) (context.Context, context.CancelFunc) {
// 	return context.WithTimeout(*parentCtx, timeout)
// }
