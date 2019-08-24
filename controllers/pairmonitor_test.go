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

func Test_IndexPage(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	pairMonitor := PairMonitor{repo: &MockRepo{}}
	pairMonitor.indexHandler(response, request)

	t.Run("returns index page", func(t *testing.T) {
		got := response.Body.String()
		want := "pair monitor"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

}

func Test_saveSession(t *testing.T) {

	t.Run("save a pairing session", func(t *testing.T) {
		pairMonitor := PairMonitor{repo: &MockRepo{}}
		response := httptest.NewRecorder()
		pairMonitor.saveSession(response, &DevPair{Dev1:"superman", Dev2:"batman"})
		var got DevPair
		json.NewDecoder(response.Body).Decode(&got)
		want := DevPair{ID: 1, Dev1: "superman", Dev2: "batman"}
		if got != want {
			t.Errorf("saveSession() = %v, want %v", got, want)
		}
	})

	t.Run("respond with error", func(t *testing.T) {
		pairMonitor := PairMonitor{repo: &MockErrorRepo{}}
		response := httptest.NewRecorder()
		pairMonitor.saveSession(response, &DevPair{Dev1:"superman", Dev2:"batman"})
		want := http.StatusInternalServerError
		got := response.Result().StatusCode
		if got != want {
			t.Errorf("status code = %v, want %v", got, want)
		}
	})
}

func Test_allSessions(t *testing.T) {

	t.Run("retrieve all sessions", func(t *testing.T) {
		pairMonitor := PairMonitor{repo: &MockRepo{}}
		response := httptest.NewRecorder()
		pairMonitor.allSessions(response)
		var want = DevPair{ID:1, Dev1:"superman", Dev2:"batman"}
		var got []DevPair
		json.NewDecoder(response.Body).Decode(&got)
		if got[0] != want {
			t.Errorf("allSessions() = %v, want %v", got, want)
		}
	})

	t.Run("respond with error", func(t *testing.T) {
		pairMonitor := PairMonitor{repo: &MockErrorRepo{}}
		response := httptest.NewRecorder()
		pairMonitor.allSessions(response)
		want := http.StatusInternalServerError
		got := response.Result().StatusCode
		if got != want {
			t.Errorf("status code = %v, want %v", got, want)
		}
	})
}
