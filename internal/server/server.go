package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zloyboy/mongo/internal/config"
	"github.com/zloyboy/mongo/internal/store"
)

type server struct {
	router *mux.Router
	store  *store.Store
}

func newServer(store *store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configRouter()
	return s
}

func Start() error {
	config := config.NewConfig()
	store, err := store.New(config)
	if err != nil {
		return err
	}
	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)
}

func (s *server) configRouter() {
	s.router.HandleFunc("/v1/start", s.handleStart()).Methods("POST")
	s.router.HandleFunc("/v1/finish", s.handleFinish()).Methods("POST")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

type request struct {
	Type string `json:"type"`
}

func (s *server) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Start(req.Type); err == nil {
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
}

func (s *server) handleFinish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
}
