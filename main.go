package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Response struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UTCTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GitHubFileURL string `json:"github_file_url"`
	GitHubRepoURL string `json:"github_repo_url"`
	StatusCode    int    `json:"-"`
}

func main() {
	http.HandleFunc("/api", apiHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	slackName := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")

	// Get current UTC time and format it
	currentTime := time.Now().UTC()
	currentTimeFormatted := currentTime.Format("2006-01-02T15:04:05Z")

	githubFileURL := "https://github.com/walerick/hngBackendFirst/blob/main/src/main.go"
	githubRepoURL := "https://github.com/walerick/hngBackendFirst.git"

	// Create the response struct
	response := Response{
		SlackName:     slackName,
		CurrentDay:    currentTime.Weekday().String(),
		UTCTime:       currentTimeFormatted,
		Track:         track,
		GitHubFileURL: githubFileURL,
		GitHubRepoURL: githubRepoURL,
		StatusCode:    http.StatusOK,
	}

	// Convert the response struct to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
