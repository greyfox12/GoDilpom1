package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/greyfox12/GoDilpom1/internal/config"

	"github.com/greyfox12/GoDilpom1/internal/app"
)

// Конфигурация по умолчанию
const (
	defServiceAddress        = "localhost:8080"
	defDSN                   = "host=localhost user=videos password=videos dbname=postgres sslmode=disable"
	defAccurualSystemAddress = "http://localhost:8090"
	defLogLevel              = "Debug"
	defAccurualTimeReset     = 120 //120 секунд - Время до сброса в БД состояния отправленных на обработку ордеров
	defIntervalAccurual      = 1   // 1 секунд - Задержка перед циклом выбора для отправки на обработку ордеров
	defTimeoutContexDB       = 10  // сек. Таймаут для контекста работы c DB
)

func initApp() {
	// Собрать конфигурацию приложения из Умолчаний, ключей и Переменнных среды
	confDef := config.Config{
		App:    config.AppConfig{Name: "YandexDiplom", LogLevel: defLogLevel},
		HTTP:   config.HTTPConfig{HostAddress: defServiceAddress, TimeoutContexDB: defTimeoutContexDB},
		MainDB: config.SQLConfig{DSN: defDSN},
		AccurualService: config.AccurualConfig{URL: defAccurualSystemAddress,
			Interval:  defIntervalAccurual,
			TimeReset: defAccurualTimeReset},
	}

	a, err := app.NewApp(&confDef)

	if err != nil {
		log.Fatal("Fail to create app: %w", err)
	}

	app.SetGlobalApp(a)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// запускаю Опрос системы начисления баллов
	if a.GetConfig().AccurualService.Interval > 0 {
		go func() {
			a.StartAccrual(ctx)
		}()
	}

	// Запускаю сервер HTTP
	if err := a.RunServer(ctx); err != nil {
		log.Fatal("start HTTP server: %w", err)
	}
}

func main() {
	initApp()

}
