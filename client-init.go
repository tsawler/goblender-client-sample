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
func ClientInit(c config.AppConfig, p *driver.DB, rep *handlers.DBRepo) {
	// c is the application config, from goblender
	app = c
	repo = rep

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
		handlers.NewPostgresqlHandlers(p, app.ServerName, app.InProduction)
	} else {
		handlers.NewMysqlHandlers(p, app.ServerName, app.InProduction)
	}

	// set different template for pages, if needed
	//repo.SetDefaultPageTemplate("client-sample.page.tmpl")

	// create client middleware
	NewClientMiddleware(app, repo)
}
