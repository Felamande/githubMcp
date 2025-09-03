package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Felamande/githubMcp/model"
	"github.com/google/go-github/v74/github"
)

// TestCommitParentComparison tests that commit parent information is correctly populated
// by comparing with direct API responses
func TestCommitParentComparison(t *testing.T) {
	// Create a test server that mimics GitHub API responses
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/repos/pytorch/pytorch/commits/e2d141dbde55c2a4370fac5165b0561b6af4798b" {
			// Return a sample commit response with parent information
			commitResponse := map[string]interface{}{
				"sha": "e2d141dbde55c2a4370fac5165b0561b6af4798b",
				"commit": map[string]interface{}{
					"message": "test commit message",
					"author": map[string]interface{}{
						"name":  "test author",
						"email": "author@example.com",
						"date":  "2023-01-01T00:00:00Z",
					},
					"committer": map[string]interface{}{
						"name":  "test committer",
						"email": "committer@example.com",
					},
				},
				"html_url": "https://github.com/pytorch/pytorch/commit/e2d141dbde55c2a4370fac5165b0561b6af4798b",
				"parents": []map[string]interface{}{
					{
						"sha": "parent1sha1234567890abcdef1234567890abcdef12",
						"url": "https://api.github.com/repos/pytorch/pytorch/commits/parent1sha1234567890abcdef1234567890abcdef12",
					},
					{
						"sha": "parent2sha1234567890abcdef1234567890abcdef12",
						"url": "https://api.github.com/repos/pytorch/pytorch/commits/parent2sha1234567890abcdef1234567890abcdef12",
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(commitResponse)
			return
		}
		http.Error(w, "Not found", http.StatusNotFound)
	}))
	defer server.Close()

	// Create client with test server URL
	githubClient := github.NewClient(nil)
	githubClient.BaseURL, _ = githubClient.BaseURL.Parse(server.URL + "/")
	client := &GithubClient{c: githubClient}

	// Test GetCommitBySHA
	commit, err := client.GetCommitBySHA(model.GetCommitBySHAOption{
		Owner:      "pytorch",
		Repository: "pytorch",
		SHA:        "e2d141dbde55c2a4370fac5165b0561b6af4798b",
	})

	if err != nil {
		t.Fatalf("GetCommitBySHA failed: %v", err)
	}

	// Check that basic commit information is populated
	if commit.SHA != "e2d141dbde55c2a4370fac5165b0561b6af4798b" {
		t.Errorf("Expected SHA %s, got %s", "e2d141dbde55c2a4370fac5165b0561b6af4798b", commit.SHA)
	}

	if commit.Message != "test commit message" {
		t.Errorf("Expected message %s, got %s", "test commit message", commit.Message)
	}

	// The issue: GetCommitBySHA doesn't populate ParentCommitHash like ListCommits does
	// This test will fail until the GetCommitBySHA function is fixed to include parent information
	if commit.ParentCommitHash == nil {
		t.Error("ParentCommitHash should not be nil")
	}

	// Expected parent SHAs from our test response
	expectedParents := []string{
		"parent1sha1234567890abcdef1234567890abcdef12",
		"parent2sha1234567890abcdef1234567890abcdef12",
	}

	if len(commit.ParentCommitHash) != len(expectedParents) {
		t.Errorf("Expected %d parents, got %d", len(expectedParents), len(commit.ParentCommitHash))
	}

	for i, expectedParent := range expectedParents {
		if i >= len(commit.ParentCommitHash) {
			break
		}
		if commit.ParentCommitHash[i] != expectedParent {
			t.Errorf("Parent %d: expected %s, got %s", i, expectedParent, commit.ParentCommitHash[i])
		}
	}
}

// TestDirectCurlComparison demonstrates how to compare with direct curl output
func TestDirectCurlComparison(t *testing.T) {
	// This test would ideally make actual HTTP requests to compare,
	// but we'll simulate it for the test environment
	t.Skip("Skipping actual HTTP request test - use this pattern for manual verification")

	// Example of how you would compare with direct curl:
	// curl -H "Accept: application/vnd.github.v3+json" https://api.github.com/repos/pytorch/pytorch/commits/e2d141dbde55c2a4370fac5165b0561b6af4798b
	
	// The response should include "parents" array with SHA values
	// Our GetCommitBySHA should return the same parent information in ParentCommitHash
}

// TestListCommitsIncludesParents confirms that ListCommits correctly includes parent information
func TestListCommitsIncludesParents(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/repos/testowner/testrepo/commits" {
			// Return a sample commits list response
			commitsResponse := []map[string]interface{}{
				{
					"sha": "testsha1234567890abcdef1234567890abcdef12",
					"commit": map[string]interface{}{
						"message": "test message",
						"author": map[string]interface{}{
							"name":  "test author",
							"email": "author@example.com",
							"date":  "2023-01-01T00:00:00Z",
						},
						"committer": map[string]interface{}{
							"name":  "test committer",
							"email": "committer@example.com",
						},
					},
					"html_url": "https://github.com/testowner/testrepo/commit/testsha1234567890abcdef1234567890abcdef12",
					"parents": []map[string]interface{}{
						{
							"sha": "parentsha1234567890abcdef1234567890abcdef12",
						},
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(commitsResponse)
			return
		}
		http.Error(w, "Not found", http.StatusNotFound)
	}))
	defer server.Close()

	// Create client with test server URL
	githubClient := github.NewClient(nil)
	githubClient.BaseURL, _ = githubClient.BaseURL.Parse(server.URL + "/")
	client := &GithubClient{c: githubClient}

	// Test ListCommits
	commits, err := client.ListCommits(model.CommitListOption{
		Owner:      "testowner",
		Repository: "testrepo",
		ResultPerpage: 10,
		Page: 1,
	})

	if err != nil {
		t.Fatalf("ListCommits failed: %v", err)
	}

	if len(commits.Commits) == 0 {
		t.Fatal("Expected at least one commit")
	}

	commit := commits.Commits[0]
	
	// ListCommits should correctly populate parent information
	if commit.ParentCommitHash == nil {
		t.Error("ParentCommitHash should not be nil in ListCommits result")
	}

	if len(commit.ParentCommitHash) != 1 {
		t.Errorf("Expected 1 parent, got %d", len(commit.ParentCommitHash))
	}

	if commit.ParentCommitHash[0] != "parentsha1234567890abcdef1234567890abcdef12" {
		t.Errorf("Expected parent SHA %s, got %s", "parentsha1234567890abcdef1234567890abcdef12", commit.ParentCommitHash[0])
	}
}