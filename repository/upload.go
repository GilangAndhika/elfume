package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/GilangAndhika/elfume/model"
)

// UploadtoGithub uploads a file to GitHub repository
func UploadtoGithub(fileName, Content string) error {
	// Get GitHub credentials from environment variables
	githubOwner := os.Getenv("GITHUB_OWNER")
	githubRepo := os.Getenv("GITHUB_REPO")
	githubToken := os.Getenv("GITHUB_TOKEN")

	// Validate GitHub credentials
	if githubOwner == "" || githubRepo == "" || githubToken == "" {
		return errors.New("GitHub connection details are missing in environment variables")
	}

	// Create GitHub API URL
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", githubOwner, githubRepo, fileName)

	// Create request payload
	uploadRequest := model.GithubUploadRequest{
		Message: "Upload file image " + fileName,
		Content: Content,
	}

	jsonData, err := json.Marshal(uploadRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+githubToken)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, _ := ioutil.ReadAll(resp.Body)

	// Fix: Correctly handle 201 Created as success
	if resp.StatusCode == http.StatusCreated {
		return nil // Success, return no error
	}

	return fmt.Errorf("failed to upload file: %s, response: %s", resp.Status, string(body))
}
