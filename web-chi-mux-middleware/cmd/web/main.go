package main

import (
	"net/http"
	"log"
	"github.com/papudatta/gowebfu/pkg/handlers"
	"github.com/papudatta/gowebfu/pkg/config"
	"github.com/papudatta/gowebfu/pkg/render"
	"github.com/alexedwards/scs/v2"
	"time"
)

const portNumber = "0.0.0.0:9090"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true for production deployment
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

//	http.HandleFunc("/", handlers.Repo.Home)
//	http.HandleFunc("/about", handlers.Repo.About)

//	_ = http.ListenAndServe(portNumber, nil)

        srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
