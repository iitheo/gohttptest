package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main(){
	//http.HandleFunc("/double", doubleHandler)
	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/doubler", doubleHandler)
	return r
}

func doubleHandler(w http.ResponseWriter, r *http.Request){
	text := r.FormValue("v")
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: " + text, http.StatusBadRequest)
		return
	}

	_,_ = fmt.Fprintln(w, v*2)
}
