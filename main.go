package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	button "github.com/livghit/templkit/components/Actions/button"
	dropdown "github.com/livghit/templkit/components/Actions/dropdown"
)

func main() {

	dropdown := dropdown.DefaultDropdown()
	o := button.ButtonOptions{
		Label: "hello",
		Class: "btn btn-primary",
	}
	button := button.DefaultButton(o)

	//mux router

	router := mux.NewRouter()

	router.Handle("/", templ.Handler(dropdown))
	router.Handle("/button", templ.Handler(button))

	log.Info("Server started at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
