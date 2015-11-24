package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type handler struct {
	couchUrl  string
	workQueue chan string
}

func NewHandler(couchUrl string, workQueue chan string) handler {
	createDatabase(couchUrl)

	return handler{couchUrl: couchUrl, workQueue: workQueue}
}

func (h handler) Handle(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)

	h.workQueue <- string(body)
}

func createDatabase(couchUrl string) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%slogs", couchUrl), nil)
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode == http.StatusNotFound {
		req, _ = http.NewRequest("PUT", fmt.Sprintf("%slogs", couchUrl), nil)
		//TODO: Handle these return args
		_, _ = (&http.Client{}).Do(req)
	}
}
