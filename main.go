package main

import (
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	println("listening on", listener.Addr().String())

	workQueue := make(chan string, 100)
	h := NewHandler("http://localhost:5984/", workQueue)

	go worker(workQueue)

	r := mux.NewRouter()
	r.Handle("/", handlers.MethodHandler{
		"POST": http.HandlerFunc(
			h.Handle,
		),
	})
	http.Handle("/", r)

	http.Serve(listener, http.DefaultServeMux)
}
