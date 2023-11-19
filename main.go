package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func wellKnownHandler(w http.ResponseWriter, r *http.Request) {

	path := os.Getenv("DID_CONFIGURATION_FILE")
	if path == "" {
		path = "/opt/well-known/did-configuration.json"
	}

	filebytes, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		http.Error(w, "failed to read configuration file", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(filebytes)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "7171"
	}

	http.HandleFunc("/.well-known/did-configuration.json", wellKnownHandler)

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
