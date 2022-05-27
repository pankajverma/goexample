package main

import (
	"fmt"
	"net/http"
	"pkg/shortner"

	"github.com/gorilla/mux"
	// "net/http"
)

func main() {

	fmt.Print("welcome to urlShortner..")
	r := mux.NewRouter()
	r.HandleFunc("/short", shortner.UrlShortner)
	r.HandleFunc("/api/v1/{code}", shortner.GetFullUrl).Methods("GET")
	// http.HandleFunc("/short", shortner.UrlShortner)
	// http.HandleFunc("/api/v1/:code", shortner.UrlShortner)

	srv := &http.Server{Handler: r,
		Addr: "127.0.0.1:8080",
	}
	srv.ListenAndServe()
	//	http.ListenAndServe(":8080", nil)

}
