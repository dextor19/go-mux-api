package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
 router := mux.NewRouter()
	router.HandleFunc("/", HomePageHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

// HomePageHandler send Hello World message to the client
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	dataToWrite := []byte("Hello World")
	w.Write(dataToWrite)
}
