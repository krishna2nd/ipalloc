package ipalloc

import (
	. "common"
	"log"
	"net/http"
	"os"
	"registry"
)

var Registry = registry.New()

func Server() {
	Registry.Load("./data/registry.txt")
	router := AppRouter()
	loggingHandler := NewApacheLoggingHandler(router, os.Stderr)
	server := &http.Server{
		Addr:    PORT,
		Handler: loggingHandler,
	}
	Log("Server starting at ", PORT)
	log.Fatal(server.ListenAndServe())
}
