package main

import (
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
    var err error
    DB, err = sqlx.Open("postgres", "postgres://life_user:life_pass123@localhost:5432/life_science_db?sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to PostgreSQL:", err)
    }

    _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        timestamp TIMESTAMP
    )`)
    if err != nil {
        log.Fatal("Failed to create table:", err)
    }
}
