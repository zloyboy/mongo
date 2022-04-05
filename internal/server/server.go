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

func isDigitLetter(s string) bool {
	for _, c := range s {
		if (c < 'a' || 'z' < c) && (c < '0' || '9' < c) {
			return false
		}
	}
	return true
}

type request struct {
	Type string `json:"type"`
}

func decodeType(r *http.Request) (etype string, res bool) {
	req := request{}

	if err := json.NewDecoder(r.Body).Decode(&req); err == nil {
		etype = req.Type
		res = isDigitLetter(etype)
	} else {
		res = false
	}

	return etype, res
}

func (s *server) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		etype, res := decodeType(r)
		if !res {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Start(etype); err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s *server) handleFinish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		etype, res := decodeType(r)
		if !res {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if finished, err := s.store.Finish(etype); err == nil {
			if finished {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
	}
}
