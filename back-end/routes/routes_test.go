package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLivenessRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello there!", w.Body.String())
}
func TestCompanyQueryString(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	companyName := "Kanelbulle"
	req, _ := http.NewRequest("GET", "/ticks?company="+companyName, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Company name provided "+companyName, w.Body.String())
}
