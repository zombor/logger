package main

import (
	"fmt"
	"net/http"
)

func setupDatabase(couchUrl string) {
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
