package main

import (
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
    c.JSON(201, event)
}
