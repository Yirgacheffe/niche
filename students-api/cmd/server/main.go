package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"students-api/internal/db"
	"time"

	internalHttp "students-api/internal/http"
	"students-api/internal/service/student"
)

func Run(wait time.Duration) error {
	log.Println("running app")

	conn, err := db.InitDatabase()
	if err != nil {
		return err
	}

	err = db.MigrateDB(conn)
	if err != nil {
		return err
	}

	service := student.NewService(conn)
	handler := internalHttp.NewHandler(service)
	handler.InitRoutes()

	srv := &http.Server{
		Addr:         ":9010",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Well, block until receive signal
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
	log.Println("shutting down...")
	os.Exit(0)

	return nil
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	err := Run(wait)
	if err != nil {
		log.Fatal("error while start app", err)
	}
}
