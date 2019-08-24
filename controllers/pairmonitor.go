package controllers

import (
	"encoding/json"
	"fmt"
	. "github.com/rajeshpg/pair-monitor-go/models"
	. "github.com/rajeshpg/pair-monitor-go/repos"
	"net/http"
)

type PairMonitor struct {
	Repo DevPairRepo
}

const jsonContentType = "application/json"

func (pairMonitor *PairMonitor) ServeHTTP(w http.ResponseWriter, r *http.Request){

	switch r.Method {
	case http.MethodPost:
		saveSession(pairMonitor.Repo, w, formValToDevPair(r))
	case http.MethodGet:
		allSessions(pairMonitor.Repo, w)
	}
}

func formValToDevPair(r *http.Request) *DevPair {
	return &DevPair{Dev1: r.FormValue("dev1"), Dev2: r.FormValue("dev2")}
}

func saveSession(repo DevPairRepo, w http.ResponseWriter, devPair *DevPair){
	_, err := repo.SaveSession(devPair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(devPair)
}

func allSessions(repo DevPairRepo, w http.ResponseWriter) {
	pairs, err := repo.AllSessions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pairs)
}
