package clienthandlers

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/middleware"
)

var app config.AppConfig
var pageModel *clientdb.PageModel

// ClientRoutes is used to handle custom routes for specific clients. Prepend some unique (and site wide) value
// to the start of each route in order to avoid clashes with pages, etc. Middleware can be applied by importing and
// using the middleware.* functions
func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	mux.Get("/custom/my-test-route", standardMiddleWare.ThenFunc(TestHandler))
	mux.Get("/custom/protected-route", standardMiddleWare.Append(middleware.Auth).ThenFunc(TestProtectedHandler))
	mux.Get("/custom/db", standardMiddleWare.Append(middleware.Auth).ThenFunc(TestDB))

	return mux, nil
}

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig) {
	app = c
	conn := app.Connections["NBTAP"]
	fmt.Print("Conn:", conn)
	pageModel = &clientdb.PageModel{DB: conn}
}
