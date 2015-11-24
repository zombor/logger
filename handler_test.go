package main

import (
	"bytes"
	"encoding/json"
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

func Test_handlerHandle_SendsReqBody_ToWorkQueue(t *testing.T) {
	jsonInput := map[string]string{
		"message": "hello world",
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(jsonInput)
	request, _ := http.NewRequest("POST", "/", body)
	response := httptest.NewRecorder()

	workQueue := make(chan string, 1)
	handler{workQueue: workQueue}.Handle(response, request)

	passedBody := <-workQueue

	assert.Equal(t, "{\"message\":\"hello world\"}\n", passedBody)
}
