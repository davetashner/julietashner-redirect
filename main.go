package main

import (
	"log"
	"net/http"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://julietashner.pixieset.com", http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", redirectHandler)

	log.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
