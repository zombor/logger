package couchdb

import (
	"bytes"
	"fmt"
	"net/http"
)

type couchDb string

func NewCouchDb(url string) couchDb {
	db := couchDb(url)
	db.setup()
	return db
}

func (db couchDb) url() string {
	return string(db)
}

func (db couchDb) setup() {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%slogs", db.url()), nil)
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err.Error())
	}

	if res.StatusCode == http.StatusNotFound {
		req, _ = http.NewRequest("PUT", fmt.Sprintf("%slogs", db.url()), nil)
		//TODO: Handle these return args
		_, _ = (&http.Client{}).Do(req)
	}

	body := bytes.NewReader([]byte(`{
		"_id": "_design/keys",
		"views": {
			"keys": {
				"map": "function(doc) { for (var thing in doc) { emit(thing,1); } }",
				"reduce": "function(key,values) { return sum(values); }"
			}
		}
	}`))
	req, _ = http.NewRequest("PUT", fmt.Sprintf("%slogs/_design/keys", db.url()), body)
	//TODO: Handle these return args
	_, _ = (&http.Client{}).Do(req)
}
