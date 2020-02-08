package clienthandlers

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/driver"
	"github.com/tsawler/goblender/pkg/handlers"
	"github.com/tsawler/goblender/pkg/repository"
	"github.com/tsawler/goblender/pkg/repository/page"
	"log"
	"net/http"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger
var pageModel repository.PageRepo
var parentDB *driver.DB

var preferenceHandlers *handlers.PreferenceDBRepo
var pageHandlers *handlers.PageDBRepo
var userHandlers *handlers.UserDBRepo
var roleHandlers *handlers.RoleDBRepo
var historyHandlers *handlers.HistoryDBRepo
var postHandlers *handlers.PostDBRepo

// ClientRoutes is used to handle custom routes for specific clients. Prepend some unique (and site wide) value
// to the start of each route in order to avoid clashes with pages, etc. Middleware can be applied by importing and
// using the middleware.* functions. All handlers are available in appHandlers, e.g. to use pageHandlers, appHandlers.PageHandlers
func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	mux.Get("/", dynamicMiddleware.ThenFunc(ShowHome))

	mux.Get("/some-handler", standardMiddleWare.ThenFunc(SomeHandler))

	// public folder
	fileServer := http.FileServer(http.Dir("./client/clienthandlers/public/"))
	mux.Get("/client/static/", http.StripPrefix("/client/static", fileServer))

	return mux, nil
}

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig, p *driver.DB) {
	app = c
	//conn := app.Connections["wheels"]
	infoLog = app.InfoLog
	errorLog = app.ErrorLog
	pageModel = page.NewSQLPageRepo(p.SQL)
	parentDB = p

	preferenceHandlers = handlers.NewPreferenceHandlers(p)
	historyHandlers = handlers.NewHistoryHandler(p)
	roleHandlers = handlers.NewRoleHandlers(p, historyHandlers)
	userHandlers = handlers.NewUserHandlers(app, p, roleHandlers)
	pageHandlers = handlers.NewPageHandler(app, p, userHandlers, preferenceHandlers)
	postHandlers = handlers.NewPostHandlers(app, p, pageHandlers)
}
