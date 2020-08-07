package main

import (
	"net/http"
)

func weightAdd(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("weight add"))
}

func weightAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("weight all"))
}

func weightEdit(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("weight edit"))
}

func weightDelete(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("weight delete"))
}
