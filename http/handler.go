package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type handler struct {
	workQueue  chan string
	repository repository
}

func NewHandler(repository repository, workQueue chan string) handler {
	return handler{repository: repository, workQueue: workQueue}
}

func (h handler) CreateLog(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)

	h.workQueue <- string(body)
}

func (h handler) Home(res http.ResponseWriter, req *http.Request) {
	t, _ := template.New("home").Parse(tplHome)

	type data struct {
		Key   string
		Value int
	}

	datas := []data{}

	for k, v := range h.repository.AllKeys() {
		datas = append(datas, data{
			Key:   k,
			Value: v,
		})
	}

	t.Execute(res, struct{ Items []data }{datas})
}
