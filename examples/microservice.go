package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var name = flag.String("service", "/admin", "/admin or /cad")
var port = flag.String("port", "3100", "Port")

func ServiceEndpoint(rw http.ResponseWriter, req *http.Request) {
	hostName, _ := os.Hostname()

	fmt.Fprintf(rw, "Service: %s, Host: %s", *name, hostName)
}

func main() {
	flag.Parse()

	log.Printf("Starting on the [port => %s, service => %s]", *port, *name)

	http.HandleFunc(*name, ServiceEndpoint)

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
