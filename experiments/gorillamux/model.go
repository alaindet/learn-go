package main

import "time"

type Todo struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Due         time.Time `json:"due"`
}
