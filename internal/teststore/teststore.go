package teststore

import "github.com/zloyboy/mongo/internal/store"

type teststore struct{}

func New() store.Store {
	return &teststore{}
}

func (s *teststore) Start(tp string) error {
	return nil
}

func (s *teststore) Finish(tp string) store.Finish {
	return store.Finish{true, nil}
}
