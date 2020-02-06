package clienthandlers

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/driver"
	"github.com/tsawler/goblender/pkg/middleware"
	"github.com/tsawler/goblender/pkg/repository"
    "github.com/tsawler/goblender/pkg/repository/page"
    "log"
)

var app config.AppConfig
var parentDB *driver.DB
var pageModel repository.PageRepo
var infoLog *log.Logger
var errorLog *log.Logger

// ClientRoutes is used to handle custom routes for specific clients. Prepend some unique (and site wide) value
// to the start of each route in order to avoid clashes with pages, etc. Middleware can be applied by importing and
// using the middleware.* functions
func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	mux.Get("/custom/my-test-route", standardMiddleWare.ThenFunc(TestHandler))
	mux.Get("/custom/protected-route", standardMiddleWare.Append(middleware.Auth).ThenFunc(TestProtectedHandler))

	return mux, nil
}

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig, p *driver.DB) {
    app = c
    infoLog = app.InfoLog
    errorLog = app.ErrorLog
    pageModel = page.NewSQLPageRepo(p.SQL)
    parentDB = p
}
