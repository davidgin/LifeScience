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
    wg.Add(2)

    go func() {
        defer wg.Done()
        result, err := DB.ExecContext(context.Background(),
            "INSERT INTO events (title, description, timestamp, processed, location) VALUES ($1, $2, $3, $4, $5) RETURNING id",
            event.Title, event.Description, event.Timestamp, false, event.Location)
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

    go func() {
        defer wg.Done()
        location, err := FetchOSMLocationWithRetry(event.Title)
        if err != nil {
            log.Println("OSM fetch failed:", err)
            return
        }
        mu.Lock()
        event.Location = location
        _, err = DB.Exec("UPDATE events SET location = $1 WHERE id = $2", location, event.ID)
        if err != nil {
            log.Println("Update location error:", err)
        }
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
