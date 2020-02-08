package clienthandlers

import (
	"net/http"
)

// JSONResponse is a generic struct to hold json responses
type JSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// ShowHome returns the home page using our local page template for the client
func ShowHome(w http.ResponseWriter, r *http.Request) {
	pageHandlers.SetDefaultPageTemplate("page.page.tmpl")
	pageHandlers.Home(w, r)
}

// other routes as necessary e.g.

// SomeHandler gets all motorcycles
func SomeHandler(w http.ResponseWriter, r *http.Request) {
	// ... do whatever
}