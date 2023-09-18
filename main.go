package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"

	button "github.com/livghit/templkit/components/Actions/button"
	"github.com/livghit/templkit/pages"
)

func main() {

	o := button.ButtonOptions{
		Label: "hello",
		Class: "btn btn-primary",
	}

	//mux router

	router := mux.NewRouter()
	index := pages.Index(o)

	router.Handle("/", templ.Handler(index))

	log.Info("Server started at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
