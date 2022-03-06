package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    InitDB()
    defer DB.Close()

    r := gin.Default()
    r.POST("/api/v1/events", CreateEvent)
    r.GET("/api/v1/events", GetEvents)
    r.Run(":8080")
}
