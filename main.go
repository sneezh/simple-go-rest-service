package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func listenAndServe(port string, router *mux.Router) *http.Server {
	srv := &http.Server{Addr: port, Handler: router}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()
	return srv
}

var config Config
var isTesting bool

func main() {
	// try to read config
	config = getConfig()

	// create db connects object
	PgConnect()
	defer db.Close()

	// add router
	router := getRouter()

	// start http server
	log.Printf("starting HTTP server")
	srv := listenAndServe(fmt.Sprintf(":%s", config["APP_PORT"]), router)

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	//log.Printf("serving for 5 seconds")
	//time.Sleep(5 * time.Second)
	log.Printf("stopping server")
	// graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	log.Printf("exit")
}
