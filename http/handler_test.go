package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zombor/logger/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func Test_handlerHandle_SendsReqBody_ToWorkQueue(t *testing.T) {
	jsonInput := map[string]string{
		"message": "hello world",
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(jsonInput)
	request, _ := http.NewRequest("POST", "/", body)
	response := httptest.NewRecorder()

	workQueue := make(chan string, 1)
	handler{workQueue: workQueue}.CreateLog(response, request)

	passedBody := <-workQueue

	assert.Equal(t, "{\"message\":\"hello world\"}\n", passedBody)
}
