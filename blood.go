package main

import (
	"net/http"
)

func bloodAdd(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("blood add"))
}

func bloodAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("blood all"))
}

func bloodDetails(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("blood details"))
}

func bloodEdit(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("blood edit"))
}

func bloodDelete(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("blood delete"))
}
