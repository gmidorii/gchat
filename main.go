package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "server port")
var historyRoot = flag.String("history", "", "history root dir")

func run() error {
	hub = NewHub(*historyRoot)
	http.HandleFunc("/chat", handler)
	return http.ListenAndServe(*addr, nil)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
