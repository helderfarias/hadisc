package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "3100", "Port")

func cadEndpoint(rw http.ResponseWriter, req *http.Request) {
	hostName, _ := os.Hostname()

	fmt.Fprintf(rw, "Service: %s, Host: %s", "cad", hostName)
	log.Printf("Service: %s, Host: %s", "cad", hostName)
}

func adminEndpoint(rw http.ResponseWriter, req *http.Request) {
	hostName, _ := os.Hostname()

	fmt.Fprintf(rw, "Service: %s, Host: %s", "admin", hostName)
	log.Printf("Service: %s, Host: %s", "admin", hostName)
}

func main() {
	flag.Parse()

	log.Printf("Starting on the [port => %s]", *port)

	http.HandleFunc("/api/cad", cadEndpoint)
	http.HandleFunc("/api/admin", adminEndpoint)

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
