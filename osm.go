package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

var OSMClient = &http.Client{Timeout: 10 * time.Second}

func FetchOSMLocationWithRetry(query string) (string, error) {
    const maxRetries = 3
    const retryDelay = 60 * time.Minute

    for attempt := 0; attempt <= maxRetries; attempt++ {
        location, err := FetchOSMLocation(query)
        if err == nil {
            return location, nil
        }
        log.Printf("OSM fetch attempt %d failed: %v", attempt+1, err)
        if attempt < maxRetries {
            log.Printf("Retrying in %d minutes...", int(retryDelay.Minutes()))
            time.Sleep(retryDelay)
        }
    }
    return "", fmt.Errorf("failed to fetch OSM location after %d retries", maxRetries)
}

func FetchOSMLocation(query string) (string, error) {
    url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", query)
    resp, err := OSMClient.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("OSM API returned status: %d", resp.StatusCode)
    }

    var results []struct {
        DisplayName string `json:"display_name"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
        return "", err
    }
    if len(results) > 0 {
        return results[0].DisplayName, nil
    }
    return "Unknown Location", nil
}
