package main

import (
	"github.com/yoyodyne-build/toolkit"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("starting server at http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/download", downloadFile)

	return mux
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	tools.DownloadStaticFile(w, r, "./files", "tipfinger.jpg", "lulzjack.jpg")
}
