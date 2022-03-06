package main

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("postgres", "postgres://life_user:life_pass123@localhost:5432/life_science_db?sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to PostgreSQL:", err)
    }
}
