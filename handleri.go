package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) DodavanjeKonfiga(w http.ResponseWriter, r *http.Request) {
	var config Config
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	s.data[id] = []*Config{&config}

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) DodavanjeGrupe(w http.ResponseWriter, r *http.Request) {
	var config Config
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	s.data[id] = []*Config{&config}

	w.WriteHeader(http.StatusCreated)
}
