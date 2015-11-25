package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zombor/logger/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func Test_setupDatabase_CreatesDatabase_IfNotExists(t *testing.T) {
	databaseExistsCalled := false
	createDatbaseCalled := false

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && r.URL.Path == "/logs" {
			databaseExistsCalled = true
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method == "PUT" && r.URL.Path == "/logs" {
			createDatbaseCalled = true
			w.WriteHeader(http.StatusCreated)
		}
	}))
	defer ts.Close()

	setupDatabase(fmt.Sprintf("%s/", ts.URL))

	assert.True(t, databaseExistsCalled)
	assert.True(t, createDatbaseCalled)
}
