package main

import (
	"log"
	"net/http"
	"time"

	"github.com/GrzegorzMika/WebApp/pkg/config"
	"github.com/GrzegorzMika/WebApp/pkg/handlers"
	"github.com/GrzegorzMika/WebApp/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PortNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = app.InProduction

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}
	log.Fatal(srv.ListenAndServe())
}
