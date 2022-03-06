package main

import (
    "context"
    "log"
    "time"
    "github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
    var event Event
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    event.Timestamp = time.Now()

    result, err := DB.ExecContext(context.Background(),
        "INSERT INTO events (title, description, timestamp) VALUES ($1, $2, $3) RETURNING id",
        event.Title, event.Description, event.Timestamp)
    if err != nil {
        log.Println("PostgreSQL error:", err)
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    var id int64
    result.Scan(&id)
    event.ID = id
    c.JSON(201, event)
}
