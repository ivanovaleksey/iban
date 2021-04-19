//go:generate swagger generate spec -o ../docs/swagger.json

package main

import (
	"context"
	"github.com/ivanovaleksey/iban/app/api"
	"github.com/ivanovaleksey/iban/app/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      api.New(),
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("listen on %s", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("server err: %v", err)
		}
	}()

	<-sign
	if err := closeSrv(srv); err != nil {
		log.Printf("close err: %v", err)
	}
}

func closeSrv(srv *http.Server) error {
	const (
		gracefulDelay   = 3 * time.Second
		gracefulTimeout = 5 * time.Second
	)

	ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancel()

	log.Println("waiting for graceful delay")
	time.Sleep(gracefulDelay)

	log.Println("shutting down")
	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("shutdown gracefully")
	return nil
}
