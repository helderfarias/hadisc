package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloServer(rw http.ResponseWriter, req *http.Request) {
	hostName, _ := os.Hostname()

	fmt.Fprintf(rw, "Hello World! ==>  Host[%s]", hostName)
}

func main() {
	http.HandleFunc("/hello", HelloServer)

	err := http.ListenAndServe(":4040", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
