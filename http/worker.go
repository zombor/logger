package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zombor/logger/Godeps/_workspace/src/github.com/nu7hatch/gouuid"
)

var worker = func(couchUrl string, ch chan string) {
	for {
		input, ok := <-ch

		if !ok {
			return
		}

		uuid, _ := uuid.NewV4()

		remoteReq, _ := http.NewRequest(
			"PUT",
			fmt.Sprintf("%slogs/%s", couchUrl, uuid.String()),
			strings.NewReader(input),
		)

		//TODO: do something with the response?
		_, err := (&http.Client{}).Do(remoteReq)

		if err != nil {
			println(err.Error())
		}
	}
}
