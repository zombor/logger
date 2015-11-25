package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/zombor/logger/Godeps/_workspace/src/github.com/gorilla/handlers"
	"github.com/zombor/logger/Godeps/_workspace/src/github.com/gorilla/mux"
)

func main() {
	listen := flag.String("listen", ":0",
		"TCP address (host:port) on which to listen for HTTP connections."+
			" Defaults to a random port."+
			" See http://golang.org/pkg/net/#Dial for examples.")
	couchUrl := flag.String("couchdb_url", "",
		"Url address for couchDb server. Must include trailing slash. Required.")

	flag.Parse()

	if *couchUrl == "" {
		panic("-couchUrl is a required flag")
	}

	listener, err := net.Listen("tcp", *listen)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	println("listening on", listener.Addr().String())

	workQueue := make(chan string, 100)
	h := NewHandler(*couchUrl, workQueue)

	go worker(*couchUrl, workQueue)

	r := mux.NewRouter()
	r.Handle("/", handlers.MethodHandler{
		"POST": http.HandlerFunc(
			h.Handle,
		),
	})
	http.Handle("/", r)

	http.Serve(listener, http.DefaultServeMux)
}
