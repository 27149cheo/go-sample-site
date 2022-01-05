package server

import (
	"context"
	"net/http"
	"time"

	"go-sample-site/pkg/core"
	"go-sample-site/pkg/log"
	"go-sample-site/pkg/server/rest"

)

func Serve(ctx context.Context) error {
	core.Start(ctx)
	defer core.Stop(ctx)

	log.Info("Start go sample site service")

	engine := rest.NewEngine()
	server := &http.Server{Addr: ":8080", Handler: engine}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Infof("Failed to stop server, error: %s\n", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Infof("Failed to start http server, error: %s\n", err)
		return err
	}

	return nil
}
