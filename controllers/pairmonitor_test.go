package controllers

import (
	"encoding/json"
	"errors"
	. "github.com/rajeshpg/pair-monitor-go/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepo struct {}

func (repo *MockRepo) SaveSession(pair *DevPair) (uint, error) {
	pair.ID = 1
	return pair.ID, nil
}

func (repo *MockRepo) AllSessions() ([]DevPair, error) {
	var sessions = []DevPair{{ID:1, Dev1:"superman", Dev2:"batman"}}
	return sessions, nil
}

type MockErrorRepo struct {}

func (repo *MockErrorRepo) SaveSession(pair *DevPair) (uint, error) {
	return 0, errors.New("database error")
}

func (repo *MockErrorRepo) AllSessions() ([]DevPair, error) {
	return nil, errors.New("database error")
}

func Test_saveSession(t *testing.T) {

	t.Run("save a pairing session", func(t *testing.T) {
		response := httptest.NewRecorder()
		saveSession(&MockRepo{}, response, &DevPair{Dev1:"superman", Dev2:"batman"})
		var got DevPair
		json.NewDecoder(response.Body).Decode(&got)
		want := DevPair{ID: 1, Dev1: "superman", Dev2: "batman"}
		if got != want {
			t.Errorf("saveSession() = %v, want %v", got, want)
		}
	})

	t.Run("respond with error", func(t *testing.T) {
		response := httptest.NewRecorder()
		saveSession(&MockErrorRepo{}, response, &DevPair{Dev1:"superman", Dev2:"batman"})
		want := http.StatusInternalServerError
		got := response.Result().StatusCode
		if got != want {
			t.Errorf("status code = %v, want %v", got, want)
		}
	})
}

func Test_allSessions(t *testing.T) {

	t.Run("retrieve all sessions", func(t *testing.T) {
		response := httptest.NewRecorder()
		allSessions(&MockRepo{}, response)
		var want = DevPair{ID:1, Dev1:"superman", Dev2:"batman"}
		var got []DevPair
		json.NewDecoder(response.Body).Decode(&got)
		if got[0] != want {
			t.Errorf("allSessions() = %v, want %v", got, want)
		}
	})

	t.Run("respond with error", func(t *testing.T) {
		response := httptest.NewRecorder()
		allSessions(&MockErrorRepo{}, response)
		want := http.StatusInternalServerError
		got := response.Result().StatusCode
		if got != want {
			t.Errorf("status code = %v, want %v", got, want)
		}
	})
}