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
		case <-time.After(2 * time.Second):
			w.Write([]byte("hello"))
		case <-r.Context().Done():
			return
		}
	})

	server := &http.Server{Addr: "localhost:3000", Handler: mux}
	errCh := make(chan error, 1)
	go func() {
		log.Println("server started")
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed %v", err)
		_ = server.Close()
	}

	if err := <-errCh; err != nil {
		log.Fatalf("server error: %v", err)
	}

	log.Println("server stopped")
}
