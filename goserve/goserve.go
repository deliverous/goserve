package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")

func main() {
	flag.Parse()
	log.Printf("About to serve '" + *root + "' on port " + *port)
	panic(http.ListenAndServe(":"+*port, handlers.LoggingHandler(os.Stdout, http.FileServer(http.Dir(*root)))))
}
