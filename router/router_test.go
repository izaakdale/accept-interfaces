package router_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/izaakdale/accept-interfaces/router"
)

type sadTestDB struct{}
type happyTestDB struct{}

func (t *sadTestDB) Ping() error {
	return errors.New("db is down")
}
func (t *happyTestDB) Ping() error {
	return nil
}

func TestRouter(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		mux := router.New(&happyTestDB{})
		rec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			t.Fail()
		}

		mux.ServeHTTP(rec, req)
		if rec.Result().StatusCode != http.StatusOK {
			t.Fail()
		}
		bodyBytes, err := io.ReadAll(rec.Body)
		if err != nil {
			t.Fail()
		}
		if string(bodyBytes) != "pong!" {
			t.Fail()
		}

		rec = httptest.NewRecorder()
		req, err = http.NewRequest(http.MethodGet, "/somethingelse", nil)
		if err != nil {
			t.Fail()
		}

		mux.ServeHTTP(rec, req)
		if rec.Result().StatusCode != http.StatusNotFound {
			t.Fail()
		}
	})

	t.Run("sad path", func(t *testing.T) {
		mux := router.New(&sadTestDB{})
		rec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			t.Fail()
		}

		mux.ServeHTTP(rec, req)

		if rec.Result().StatusCode != http.StatusInternalServerError {
			t.Error("failing DB does not cause http failure")
		}
	})
}
