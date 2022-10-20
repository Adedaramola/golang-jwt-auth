package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adedaramola/golang-jwt-auth/auth"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/ping", PingServer)
	router.POST("/auth", LoginUser)
	router.GET("/protected", auth.Authenticate(Protected))

	server := &http.Server{
		Addr:         ":5001",
		Handler:      router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Println("starting server")

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server")
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	log.Println("server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("failed to shutdown server properly", err)
	}

	log.Println("server exited gracefully")
}
