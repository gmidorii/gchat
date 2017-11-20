package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "server port")

func run() error {
	http.HandleFunc("/chat", handler)
	return http.ListenAndServe(*addr, nil)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
