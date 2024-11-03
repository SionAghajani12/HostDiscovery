package main

import (
	"log"
	"net/http"
)

func main() {
	// Set the file paths to the executables
	linuxExecutablePath := "./ping_sweep_arm64"
	macExecutablePath := "./ping_sweep_arm64_mac"
	windowsExecutablePath := "./ping_sweep_arm64_windows.exe"

	// Serve the index.html file and the executables
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ping_sweep_arm64", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, linuxExecutablePath)
	})

	http.HandleFunc("/ping_sweep_arm64_mac", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, macExecutablePath)
	})

	http.HandleFunc("/ping_sweep_arm64_windows.exe", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, windowsExecutablePath)
	})

	// Start the HTTP server
	port := "8080"
	log.Printf("Starting server on :%s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
