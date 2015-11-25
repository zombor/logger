package main

import (
	"io/ioutil"
	"net/http"
)

type handler struct {
	couchUrl  string
	workQueue chan string
}

func NewHandler(couchUrl string, workQueue chan string) handler {
	return handler{couchUrl: couchUrl, workQueue: workQueue}
}

func (h handler) CreateLog(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)

	h.workQueue <- string(body)
}

func (h handler) Home(res http.ResponseWriter, req *http.Request) {
}
