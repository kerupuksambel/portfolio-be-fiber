package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kerupuksambel/portfolio-be/app"
)

func TestEndpoints(t *testing.T) {
	server := app.Init()

	endpoints := []string{
		"/api/works",
		"/api/projects",
		// "/api/notfound",
	}

	for _, endpoint := range endpoints {
		req := httptest.NewRequest(http.MethodGet, endpoint, nil)
		resp, _ := server.Test(req)

		assert.Equal(t, resp.StatusCode, http.StatusOK)
	}
}
