package clienthandlers

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

// ClientRoutes is used to handle custom routes for specific clients. Prepend some unique (and site wide) value
// to the start of each route in order to avoid clashes with pages, etc. Middleware can be applied by importing and
// using the middleware.* functions. All handlers are available in appHandlers, e.g. to use pageHandlers, appHandlers.PageHandlers
func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	// we can override routes in goblender, if we wish, e.g.
	//mux.Get("/", dynamicMiddleware.ThenFunc(pageHandlers.Home))

	mux.Get("/client/some-handler", standardMiddleWare.ThenFunc(SomeHandler))

	// public folder
	fileServer := http.FileServer(http.Dir("./client/clienthandlers/public/"))
	mux.Get("/client/static/", http.StripPrefix("/client/static", fileServer))

	return mux, nil
}
