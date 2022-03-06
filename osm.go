package main

import (
    "net/http"
)

var OSMClient = &http.Client{Timeout: 10 * time.Second}
