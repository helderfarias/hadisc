ackage main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "3100", "Port")

func Cad(rw http.ResponseWriter, req *http.Request) {
	hostName, _ := os.Hostname()

	fmt.Fprintf(rw, "Hello World, Cad! ==>  Host[%s]", hostName)
}

func main() {
	flag.Parse()

	log.Printf("Starting on the port => %s", *port)

	http.HandleFunc("/api/v2/cad", Cad)

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
