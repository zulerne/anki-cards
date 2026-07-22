package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			return
		case <-time.After(5 * time.Second):
			w.Write([]byte("OK"))
		}
	})

	srv := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	errCh := make(chan error, 1)

	go func() {
		log.Println("server started")
		err := srv.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
				return
			}
			errCh <- nil
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("graceful shutdown failed")
		srv.Close()
	}

	if err := <-errCh; err != nil {
		log.Printf("server error: %v", err)
	}

	log.Println("graceful shutdown")
}
