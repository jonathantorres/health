package main

import (
	"net/http"
)

func bloodAdd(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Blood Pressure Add Reading"
	if err := renderView("views/blood/add.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Blood Pressure Readings"
	if err := renderView("views/blood/all.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodDetails(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Blood Pressure Reading Details"
	if err := renderView("views/blood/details.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodEdit(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	layoutData.PageTitle = "Health - Edit Blood Pressure Reading"
	if err := renderView("views/blood/edit.html", res); err != nil {
		serveViewError(res, err)
	}
}

func bloodDelete(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")
	res.Write([]byte("blood delete"))
}
