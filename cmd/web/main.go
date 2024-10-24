package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/sigivooim/bookings/internal/config"
	"github.com/sigivooim/bookings/internal/handlers"
	"github.com/sigivooim/bookings/internal/render"
	"log"
	"net/http"
	"time"
)

// App Config
var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

// main is the main application function
func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// create the templatecache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache: ", err)
	}

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// add template cache to the app config
	app.TemplateCache = tc
	app.UseCache = false

	// add app config to render package
	render.NewRenderer(&app)

	fmt.Println("Starting application on port " + portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
