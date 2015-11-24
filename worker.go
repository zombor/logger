package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

var worker = func(ch chan string) {
	for {
		input, ok := <-ch

		if !ok {
			return
		}

		uuid, _ := exec.Command("uuidgen").Output()
		sUuid := strings.TrimSpace(string(uuid))

		remoteReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("http://localhost:5984/logs/%s", sUuid),
			strings.NewReader(input),
		)

		//TODO: do something with the response?
		_, err := (&http.Client{}).Do(remoteReq)

		if err != nil {
			println(err.Error())
		}
	}
}
