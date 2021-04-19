package api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPI_ValidateIBAN(t *testing.T) {
	t.Run("with empty parameter", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/iban/validate", nil)
		require.NoError(t, err)

		fx.handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		expectedBody := `{"error":"bad request","debug":"empty parameter"}`
		assert.Equal(t, expectedBody, strings.TrimRight(rr.Body.String(), "\n"))
	})

	t.Run("with invalid format", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/iban/validate?iban=BRAA00360305000010009795493P1", nil)
		require.NoError(t, err)

		fx.handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		expectedBody := `{"error":"bad request","debug":"iban: invalid format"}`
		assert.Equal(t, expectedBody, strings.TrimRight(rr.Body.String(), "\n"))
	})

	t.Run("with invalid checksum", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/iban/validate?iban=BR9700360305000010009795493P2", nil)
		require.NoError(t, err)

		fx.handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		expectedBody := `{"error":"bad request","debug":"iban: invalid checksum"}`
		assert.Equal(t, expectedBody, strings.TrimRight(rr.Body.String(), "\n"))
	})

	t.Run("with valid input", func(t *testing.T) {
		fx := newFixture(t)
		defer fx.Finish()

		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/iban/validate?iban=BR9700360305000010009795493P1", nil)
		require.NoError(t, err)

		fx.handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		assert.Empty(t, rr.Body.String())
	})
}

type fixture struct {
	t   *testing.T
	ctx context.Context

	api     *API
	handler http.HandlerFunc
}

func newFixture(t *testing.T) *fixture {
	fx := &fixture{
		t:   t,
		ctx: context.Background(),
		api: New(),
	}
	fx.handler = fx.api.ValidateIBAN()
	return fx
}

func (fx *fixture) Finish() {
}
