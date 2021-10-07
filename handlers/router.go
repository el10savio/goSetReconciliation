package handlers

import (
	"fmt"
	"net/http"

	"set"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	// Set
	Set set.Set
)

func init() {
	Set = set.Initialize()
}

// Route defines the Mux
// router individual route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a collection
// of individual Routes
var Routes = []Route{
	{"/", "GET", Index},
	{"/set/list", "GET", List},
	{"/set/add", "POST", Add},
}

// Index is the handler for the path "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Set Node\n")
}

// Logger is the middleware to
// log the incoming request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")
		next.ServeHTTP(w, r)
	})
}

// Router returns a mux router
func Router() *mux.Router {
	// Initialize Router
	router := mux.NewRouter()

	// Instantiate Routes
	for _, route := range Routes {
		router.HandleFunc(
			route.Path,
			route.Handler,
		).Methods(route.Method)
	}

	// Enable Router Logger
	router.Use(Logger)

	return router
}
