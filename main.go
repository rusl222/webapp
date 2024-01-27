package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/rusl222/webapp/model"
	"github.com/rusl222/webapp/server"
	"github.com/rusl222/webapp/view"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {

	flag.Parse()
	hub := server.NewHub()

	hub.AddCallback("printEcho", model.PrintEcho)

	go hub.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHome(view.GetHomePage(), w, r)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWs(hub, w, r)
	})

	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
