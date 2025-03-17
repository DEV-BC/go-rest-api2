package http

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

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	h.Router.HandleFunc("/api/v1/comments", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/v1/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comments/{id}", h.DeleteComment).Methods("DELETE")

}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("shutting down server gracefully")

	return nil
}

/*
	Gracefully shutting down server by wrapping in a goroutine and creating a channel to listen for an os.Interrupt signal
	from the signal package. Once the channel receives the os.Interrupt, a context will be created with a timeout of 15 secs.
	Once 15 secs is up, cancel will be called and then the server will gracefully shutdown after all connections return to idle state
	meaning that processes will be able to finish before its shutdown
*/
