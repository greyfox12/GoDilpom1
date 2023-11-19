// Старт сервиса работы с системой расчерта бонусов
package app

import (
	"context"
	"fmt"
	"time"

	"github.com/greyfox12/GoDilpom1/internal/closer"
	v1 "github.com/greyfox12/GoDilpom1/internal/handler/http/api/v1"
)

func (a *App) startAccrual(ctx context.Context) error {

	if a.cfg.AccurualService.Interval > 0 {
		go func() {
			a.startAccrual(ctx)
		}()

	}
	return nil
}

func (a *App) StartAccrual(ctx context.Context) error {

	var closer = &closer.Closer{}
	handler := v1.NewHandler(a.c.GetUseCase(), a.logger, a.cfg)

	ticker := time.NewTicker(time.Second * time.Duration(a.cfg.AccurualService.Interval))
	stop := make(chan bool)

	closer.Add(func(ctx context.Context) error {
		time.Sleep(1 * time.Second)
		stop <- true
		return nil
	})

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				handler.ExecAccrual(ctx)
			case <-stop:
				return
			}
		}
	}()

	a.logger.Info(fmt.Sprintf("start RunAccurual on %s", a.cfg.AccurualService.Url))
	<-ctx.Done()

	a.logger.Info("shutting down RunAccurual gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := closer.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer: %v", err)
	}

	return nil

}
