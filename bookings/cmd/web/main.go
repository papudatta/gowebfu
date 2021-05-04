package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/papudatta/bookings/internal/config"
	"github.com/papudatta/bookings/internal/driver"
	"github.com/papudatta/bookings/internal/handlers"
	"github.com/papudatta/bookings/internal/helpers"
	"github.com/papudatta/bookings/internal/models"
	"github.com/papudatta/bookings/internal/render"
)

const portNumber = "0.0.0.0:9090"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to our db ...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=StrongAdminP@ssw0rd")
	if err != nil {
		log.Fatal("cannot connect to DB. This is fatal!")
	}

	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}

// func main_1() {

// 	// related to session
// 	gob.Register(models.Reservation{})
// 	gob.Register(models.User{})
// 	gob.Register(models.Room{})
// 	gob.Register(models.Restriction{})

// 	// change this to true for production deployment
// 	app.InProduction = false

// 	session = scs.New()
// 	session.Lifetime = 24 * time.Hour
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteLaxMode
// 	session.Cookie.Secure = app.InProduction

// 	app.Session = session

// 	// connect to our postgresdb
// 	log.Println("Connecting to db ...")
// 	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=StrongAdminP@ssw0rd")
// 	if err != nil {
//                 log.Fatal("cannot connect to DB. This is fatal!")
//         }

// 	tc, err := render.CreateTemplateCache()
// 	if err != nil {
// 		log.Fatal("cannot create template cache")
// 	}

// 	app.TemplateCache = tc
// 	app.UseCache = false

// 	repo := handlers.NewRepo(&app, db)
// 	handlers.NewHandlers(repo)

// 	render.NewRenderer(&app)
// 	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

// //	http.HandleFunc("/", handlers.Repo.Home)
// //	http.HandleFunc("/about", handlers.Repo.About)

// //	_ = http.ListenAndServe(portNumber, nil)

//         srv := &http.Server {
// 		Addr: portNumber,
// 		Handler: routes(&app),
// 	}

// 	err = srv.ListenAndServe()
// 	if err != nil{
// 	        log.Fatal(err)
//         }

// }
