package app

import (
	"context"
	"fmt"

	"net/http"
	"time"

	"github.com/greyfox12/GoDilpom1/internal/closer"
	httpServ "github.com/greyfox12/GoDilpom1/internal/handler/http"
	v1 "github.com/greyfox12/GoDilpom1/internal/handler/http/api/v1"
)

func (a *App) RunServer(ctx context.Context) error {

	var closer = &closer.Closer{}

	handler := v1.NewHandler(a.c.GetUseCase(), a.logger, a.cfg)

	router := httpServ.NewRouter()

	router.WithHandler(handler, a.logger, a.cfg)

	srv := httpServ.NewServer(a.cfg.HTTP)

	srv.RegisterRoutes(router)

	closer.Add(func(ctx context.Context) error { return srv.Stop() })

	closer.Add(func(ctx context.Context) error {
		time.Sleep(3 * time.Second)
		return nil
	})

	go func() {
		if err := srv.ListenServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal(fmt.Sprintf("listen and serve: %v", err))
		}
	}()

	a.logger.Info(fmt.Sprintf("listening on %s", a.cfg.HTTP.HostAddress))
	<-ctx.Done()

	a.logger.Info("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := closer.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil

}
