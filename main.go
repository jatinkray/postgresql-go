package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	// Import the `pq` package with a preceding underscore since it is imported as a side
	// effect. The `pq` package is a GO Postgres driver for the `database/sql` package.
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func newRouter() *mux.Router {
	// Declare a router
	r := mux.NewRouter()
	// Declare static file directory
	// staticFileDirectory := http.Dir("./static/")
	// Create static file server for our static files, i.e., .html, .css, etc

	box := packr.NewBox("./static/")
	staticFileServer := http.FileServer(box)

	// Create file handler. Although the static files are placed inside `./static/` folder
	// in our local directory, it is served at the root (i.e., http://localhost:8080/)
	// when browsed in a browser. Hence, we need `http.StripPrefix` function to change the
	// file serve path.
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	// Add static file handler to our router
	r.Handle("/", staticFileHandler).Methods("GET")
	// Add handler for `get` and `post` people functions
	r.HandleFunc("/person", getPersonHandler).Methods("GET")
	r.HandleFunc("/server", getServerHandler).Methods("GET")
	r.HandleFunc("/person", createPersonHandler).Methods("POST")

	return r
}
func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}

	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

func main() {

	initDb()

	store = &dbStore{db: db}

	// Create router
	r := newRouter()

	// Listen to the port. Go server's default port is 8080.
	http.ListenAndServe(":8080", r)
}
