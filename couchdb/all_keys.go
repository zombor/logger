package couchdb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (db couchDb) AllKeys() map[string]int {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%slogs/_design/keys/_view/keys?group=true", db.url()),
		nil,
	)
	res, _ := (&http.Client{}).Do(req)

	var resBody struct {
		Rows []struct {
			Key   string
			Value int
		}
	}
	json.NewDecoder(res.Body).Decode(&resBody)

	data := make(map[string]int)
	for _, row := range resBody.Rows {
		data[row.Key] = row.Value
	}

	return data
}
