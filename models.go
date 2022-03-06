package main

import "time"

type Event struct {
    ID          int64     `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    Timestamp   time.Time `json:"timestamp" db:"timestamp"`
    Processed   bool      `json:"processed" db:"processed"`
}
