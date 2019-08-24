package controllers

import (
	"encoding/json"
	"fmt"
	. "github.com/rajeshpg/pair-monitor-go/models"
	. "github.com/rajeshpg/pair-monitor-go/repos"
	"net/http"
)

type PairMonitor struct {
	repo DevPairRepo
	http.Handler
}

const jsonContentType = "application/json"

func NewPairMonitor(repo *DevPairDao) *PairMonitor {

	pairMonitor := &PairMonitor{repo: repo}

	router := http.NewServeMux()
	router.HandleFunc("/", pairMonitor.indexHandler)
	router.HandleFunc("/sessions", pairMonitor.sessionsHandler)

	pairMonitor.Handler = router
	return pairMonitor
}

func (pairMonitor *PairMonitor) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "pair monitor")
}

func (pairMonitor *PairMonitor) sessionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		pairMonitor.saveSession(w, formValToDevPair(r))
	case http.MethodGet:
		pairMonitor.allSessions(w)
	}
}

func formValToDevPair(r *http.Request) *DevPair {
	return &DevPair{Dev1: r.FormValue("dev1"), Dev2: r.FormValue("dev2")}
}

func (pairMonitor *PairMonitor) saveSession(w http.ResponseWriter, devPair *DevPair) {
	_, err := pairMonitor.repo.SaveSession(devPair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(devPair)
}

func (pairMonitor *PairMonitor) allSessions(w http.ResponseWriter) {
	pairs, err := pairMonitor.repo.AllSessions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pairs)
}
