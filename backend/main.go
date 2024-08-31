package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/browser"
	"github.com/retrokyo/basedchat/internal/spahandler"
)

func main() {
	router := mux.NewRouter()

	spa := spahandler.SpaHandler{StaticPath: "build", IndexPath: "index.html"}

	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	browser.OpenURL("http://localhost:3000")

	log.Fatal(srv.Serve(listener))
}
