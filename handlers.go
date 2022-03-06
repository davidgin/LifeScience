package main

import (
    "context"
    "log"
    "sync"
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

    var mu sync.Mutex
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        result, err := DB.ExecContext(context.Background(),
            "INSERT INTO events (title, description, timestamp) VALUES ($1, $2, $3) RETURNING id",
            event.Title, event.Description, event.Timestamp)
        if err != nil {
            log.Println("PostgreSQL error:", err)
            return
        }
        var id int64
        result.Scan(&id)
        mu.Lock()
        event.ID = id
        mu.Unlock()
    }()

    wg.Wait()
    c.JSON(201, event)
}

func GetEvents(c *gin.Context) {
    var events []Event
    err := DB.Select(&events, "SELECT * FROM events ORDER BY timestamp DESC")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, events)
}
