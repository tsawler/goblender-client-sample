package clienthandlers

import (
	"net/http"
)

// JSONResponse is a generic struct to hold json responses
type JSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// ShowHome is a sample handler which returns the home page using our local page template for the client,
// and is called from client-routes.go using a route that overrides the one in goblender. This allows us
// to build custom functionality without having to use non-standard routes.
//
// Note that the template below doesn't actually exist in this repo ,and this handler is never really called.
func ShowHome(w http.ResponseWriter, r *http.Request) {
	pageHandlers.SetDefaultPageTemplate("client-page.page.tmpl")
	pageHandlers.Home(w, r)
}


// SomeHandler is an example handler
func SomeHandler(w http.ResponseWriter, r *http.Request) {
	// ... do whatever
}