package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewHandler_CreatesDatabase_IfNotExists(t *testing.T) {
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

	NewHandler(fmt.Sprintf("%s/", ts.URL), make(chan string, 0))

	assert.True(t, databaseExistsCalled)
	assert.True(t, createDatbaseCalled)
}
