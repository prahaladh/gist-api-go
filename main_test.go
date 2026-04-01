	package main

	import (
		"net/http"
		"net/http/httptest"
		"testing"
		"github.com/stretchr/testify/assert"
	)

	func TestGetGists(t *testing.T) {
		// Initialize the router
		router := SetupRouter()

		// Record the response
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/octocat", nil)
		
		// Execute the request
		router.ServeHTTP(w, req)

		// Validate results
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "https://gist.github.com/")
	}