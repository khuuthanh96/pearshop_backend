package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	appConfig "pearshop_backend/app/config"
	"pearshop_backend/app/delivery/http/payload"
	_ "pearshop_backend/app/delivery/http/payload"
	"pearshop_backend/app/delivery/http/routes"
	"pearshop_backend/pkg/hashid"
	"pearshop_backend/pkg/log"
)

func main() {
	cfg, err := appConfig.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("load app config: %w", err))
	}

	// initialization
	initDBConnection(cfg.MySQL)
	payload.LoadPayloadConfigFile("assets/payload_field_config.yaml")
	if err := hashid.InitIDHasher(cfg.IDHasher.MinLength, cfg.IDHasher.Salt); err != nil {
		panic(fmt.Errorf("init id hasher: %w", err))
	}

	stopSignal := make(chan os.Signal, 1)

	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	// #nosec
	s := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: routes.Handler(&cfg),
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Errorln("ListenAndServe fail")

			panic(err)
		}
	}()

	log.WithField("port", cfg.Port).Debugln("server started")

	<-stopSignal

	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Info("server stopped")
}
