package clienthandlers

import (
	"github.com/tsawler/goblender/pkg/helpers"
	"github.com/tsawler/goblender/pkg/templates"
	"net/http"
)

// SomeHandler is an example handler
func SomeHandler(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r, "client-sample.page.tmpl", &templates.TemplateData{})
}

// The next func is a sample handler which returns the home page using our local page template for the client,
// and is called from client-routes.go using a route that overrides the one in goblender. This allows us
// to build custom functionality without having to use non-standard routes.
// **Note** that the template below *doesn't actually exist* in this repo ,and this handler is never really called.

// ShowHome shows home page using custom template, and overrides the handler in default goblender application
//func ShowHome(w http.ResponseWriter, r *http.Request) {
//	pageHandlers.SetDefaultPageTemplate("custom-home--page.page.tmpl")
//	pageHandlers.Home(w, r)
//}


