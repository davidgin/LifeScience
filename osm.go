package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

var OSMClient = &http.Client{Timeout: 10 * time.Second}

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
