package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/zombor/logger/Godeps/_workspace/src/github.com/nu7hatch/gouuid"
)

var worker = func(couchUrl string, ch chan string) {
	for {
		input, ok := <-ch

		if !ok {
			return
		}

		var inputJson map[string]interface{}
		err := json.NewDecoder(strings.NewReader(input)).Decode(&inputJson)

		if err == nil {
			inputJson["created_at"] = time.Now().UTC()

			output, _ := json.Marshal(inputJson)

			uuid, _ := uuid.NewV4()

			remoteReq, _ := http.NewRequest(
				"PUT",
				fmt.Sprintf("%slogs/%s", couchUrl, uuid.String()),
				strings.NewReader(string(output)),
			)

			//TODO: do something with the response?
			_, err := (&http.Client{}).Do(remoteReq)

			if err != nil {
				println(err.Error())
			}
		} else {
			println(err.Error())
		}
	}
}
