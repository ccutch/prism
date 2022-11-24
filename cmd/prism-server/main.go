package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"prism.made-by-connor.com/prism"
)

func main() {
	// Binding CLI methods to api routes via mux router
	mux := http.NewServeMux()
	mux.HandleFunc("/deploy", deployHandler)
	// Running server on given port
	port := os.Getenv("PORT")
	if port == "" {
		port = "4040"
	}
	fmt.Println("Running server on port " + port)
	panic(http.ListenAndServe("0.0.0.0:"+port, mux))
}

func deployHandler(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Query().Get("source")
	if err := prism.Deploy(src); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&err)
	}
	fmt.Fprintln(w, "Successfully deployed "+src)
}
