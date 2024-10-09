package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/israelluze/go-hexagonal/adapters/web/server/handler"
	"github.com/israelluze/go-hexagonal/application"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	router := mux.NewRouter()
	negroniHandler := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(router, negroniHandler, w.Service)
	http.Handle("/", router)

	// // Use the router and negroniHandler
	// router.Handle("/products/{id}", negroniHandler.With(
	// 	negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("Hello World"))
	// 	})),
	// )).Methods("GET")

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux, // Use the router here
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
