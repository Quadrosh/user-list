package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	api "github.com/quadrosh/user-list/api"
	rp "github.com/quadrosh/user-list/repository/postgres"

	"github.com/quadrosh/user-list/user"
)

// App config
type App struct {
	Router *mux.Router
	Flags  map[string]string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := run(); err != nil {
		log.Panic(err)
	}
}

// Initialize the app
func (a *App) Initialize() {

	flags := make(map[string]string)

	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "", "Database name")
	dbUser := flag.String("dbuser", "", "Database user")
	dbPass := flag.String("dbpass", "", "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	port := flag.String("port", "8080", "Application serve port")

	flag.Parse()
	flags["dbName"] = *dbName
	flags["dbUser"] = *dbUser
	flags["dbPass"] = *dbPass
	flags["dbHost"] = *dbHost
	flags["dbPort"] = *dbPort
	flags["port"] = *port

	a.Router = mux.NewRouter()
	a.Flags = flags
}

func run() error {
	a := App{}

	a.Initialize()

	repo, err := rp.ConnectPostgres(a.Flags["dbHost"], a.Flags["dbPort"], a.Flags["dbUser"], a.Flags["dbName"], a.Flags["dbPass"])
	if err != nil {
		log.Fatal(err)
	}

	service := user.NewService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/create", handler.Create)
	r.Post("/find", handler.Find)
	r.Post("/filter", handler.Filter)

	errs := make(chan error, 2)

	go func() {
		fmt.Printf("listening on port :%s\n", a.Flags["port"])
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", a.Flags["port"]), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

	return nil
}
