package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

var worker = func(couchUrl string, ch chan string) {
	for {
		input, ok := <-ch

		if !ok {
			return
		}

		uuid, _ := exec.Command("uuidgen").Output()
		sUuid := strings.TrimSpace(string(uuid))

		remoteReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("%slogs/%s", couchUrl, sUuid),
			strings.NewReader(input),
		)

		//TODO: do something with the response?
		_, err := (&http.Client{}).Do(remoteReq)

		if err != nil {
			println(err.Error())
		}
	}
}
