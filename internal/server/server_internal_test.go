package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zloyboy/mongo/internal/teststore"
)

type testEvent struct {
	Type   string
	Status int
}

func TestSever_HandleStartOk(t *testing.T) {
	b := &bytes.Buffer{}
	event := testEvent{Type: "e0", Status: 0}
	json.NewEncoder(b).Encode(event)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/start", b)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestSever_HandleStartBad(t *testing.T) {
	b := &bytes.Buffer{}
	event := testEvent{Type: "E0", Status: 0}
	json.NewEncoder(b).Encode(event)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/start", b)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusBadRequest)
}

func TestSever_HandleFinishOk(t *testing.T) {
	b := &bytes.Buffer{}
	event := testEvent{Type: "e0", Status: 1}
	json.NewEncoder(b).Encode(event)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/finish", b)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestSever_HandleFinishBad(t *testing.T) {
	b := &bytes.Buffer{}
	event := testEvent{Type: "E0", Status: 0}
	json.NewEncoder(b).Encode(event)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/finish", b)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusBadRequest)
}
