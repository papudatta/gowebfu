package main

import (
	"net/http"
	"log"
	"github.com/papudatta/webfu/pkg/handlers"
	"github.com/papudatta/webfu/pkg/config"
	"github.com/papudatta/webfu/pkg/render"
)

const portNumber = "0.0.0.0:9090"


func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(portNumber, nil)
}
