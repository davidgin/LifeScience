package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.POST("/api/v1/events", CreateEvent)
    r.Run(":8080")
}
