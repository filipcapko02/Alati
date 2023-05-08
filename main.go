package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
<<<<<<< HEAD
	fmt.Println("Hello World")
	fmt.Println("Momcilo Sokic SR552021")
=======
	quit := make(chan os.Signal)
	service := NewService()

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/config/{id}", service.DodavanjeKonfiga).Methods("POST")
	router.HandleFunc("/config-group/{id}", service.DodavanjeGrupe).Methods("POST")

	//log.Fatal(http.ListenAndServe(":8080", router))

	// start server
	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")

}

type Config struct {
	Entries map[string]string `json:"entries"`
}

type Service struct {
	data map[string][]*Config `json:"data"`
}

func NewService() *Service {
	return &Service{
		data: make(map[string][]*Config),
	}
>>>>>>> a2e1e33 (Napravljeni Servis i Config struktura i dodavanje.)
}


