package main

import (
	"net/http"
)


func main(){

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello Universe"))
}
