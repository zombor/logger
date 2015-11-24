package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
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
	//var inputJson map[string]string
	//_ = json.NewDecoder(req.Body).Decode(&inputJson)

	body, _ := ioutil.ReadAll(req.Body)

	h.workQueue <- string(body)
}

var worker = func(ch chan string) {
	for {
		input, ok := <-ch

		if !ok {
			return
		}

		uuid, _ := exec.Command("uuidgen").Output()

		remoteReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("http://localhost:5984/logs/%s", uuid),
			strings.NewReader(input),
		)

		remoteRes, err := (&http.Client{}).Do(remoteReq)

		if err != nil {
			println(err.Error())
		}

		fmt.Printf("%#v\n", remoteRes)
	}
}

func createDatabase(couchUrl string) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%slogs", couchUrl), nil)
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode == http.StatusNotFound {
		req, _ = http.NewRequest("PUT", fmt.Sprintf("%slogs", couchUrl), nil)
		_, _ = (&http.Client{}).Do(req)
	}
}
