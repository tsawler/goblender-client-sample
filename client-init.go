package clienthandlers

import (
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/driver"
	"github.com/tsawler/goblender/pkg/handlers"
	"log"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger
var parentDB *driver.DB

var repo *handlers.DBRepo

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig, p *driver.DB, r *handlers.DBRepo) {
	// c is the application config, from goblender
	app = c
	repo = r

	// if we have additional databases (external to this application) we set the connection here
	// The connection is specified in goBlender preferences
	//conn := app.AlternateConnection

	// loggers
	infoLog = app.InfoLog
	errorLog = app.ErrorLog

	// in case we need it, we get the db connection from goblender and save it in a variable
	parentDB = p

	// if we want a local model, eg one to hit pages in goblender's db:
	//pageModel = page.NewSQLPageRepo(p.SQL)

	// we can access handlers from goblender, but need to initialize them first
	if app.Database == "postgresql" {
		//preferenceHandlers = handlers.NewPostgresPreferenceHandlers(p)
		//historyHandlers = handlers.NewPostgresHistoryHandler(p)
		//roleHandlers = handlers.NewPostgresRoleHandlers(p, historyHandlers)
		//userHandlers = handlers.NewPostgresUserHandlers(app, p, roleHandlers)
		//pageHandlers = handlers.NewPostgresPageHandler(app, p, userHandlers, preferenceHandlers, backupRepo)
		//postHandlers = handlers.NewPostgresPostHandlers(app, p, pageHandlers)
		handlers.NewPostgresqlHandlers(p, app.ServerName, app.InProduction)
	} else {
		//preferenceHandlers = handlers.NewPreferenceHandlers(p)
		//historyHandlers = handlers.NewHistoryHandler(p)
		//roleHandlers = handlers.NewRoleHandlers(p, historyHandlers)
		//userHandlers = handlers.NewUserHandlers(app, p, roleHandlers)
		//pageHandlers = handlers.NewPageHandler(app, p, userHandlers, preferenceHandlers, backupRepo)
		//postHandlers = handlers.NewPostHandlers(app, p, pageHandlers)
		handlers.NewMysqlHandlers(p, app.ServerName, app.InProduction)
	}

	// create client middleware
	NewClientMiddleware(app, repo)
}
