package couchdb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zombor/logger/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func Test_AllKeys_ReturnsData(t *testing.T) {
	dbCalled := false

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbCalled = true
		assert.Equal(t, "/logs/_design/keys/_view/keys", r.URL.Path)
		assert.Equal(t, "group=true", r.URL.RawQuery)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"rows":[{"key":"foo","value":1}]}`)
	}))
	defer ts.Close()

	db := couchDb(ts.URL + "/")
	res := db.AllKeys()

	assert.True(t, dbCalled)
	assert.Equal(t, map[string]int{"foo": 1}, res)
}
