package config

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type AppConfig struct {
	Name     string
	LogLevel string // Уровень логирования
}

type HTTPConfig struct {
	HostAddress     string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	TimeoutContexDB int // сек. Таймаут для контекста работы c DB

}

type AccurualConfig struct {
	Url       string
	TimeReset int // Время после которого сбрасывается запрос к системе начисления баллов
	Interval  int // Интервал в секудах опроса системы начисления баллов

}

type SQLConfig struct {
	DSN          string
	Timeout      time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Config struct {
	App             AppConfig      // Приложение
	HTTP            HTTPConfig     // HTTP service
	MainDB          SQLConfig      // DB
	AccurualService AccurualConfig // Out service
}

func NewConfig(sp *Config) (Config, error) {

	var ok bool
	var tStr string

	conf := Config{}

	// Копирую параметры в конфигурацию, пока их не настраиваем отдельно
	conf.AccurualService.TimeReset = sp.AccurualService.TimeReset
	conf.AccurualService.Interval = sp.AccurualService.Interval
	conf.HTTP.TimeoutContexDB = sp.HTTP.TimeoutContexDB

	if conf.HTTP.HostAddress, ok = os.LookupEnv("RUN_ADDRESS"); !ok {
		conf.HTTP.HostAddress = sp.HTTP.HostAddress
	}

	if conf.AccurualService.Url, ok = os.LookupEnv("ACCRUAL_SYSTEM_ADDRESS"); !ok {
		conf.AccurualService.Url = sp.AccurualService.Url
	}

	if conf.App.LogLevel, ok = os.LookupEnv("LOGING_LEVEL"); !ok {
		conf.App.LogLevel = sp.App.LogLevel
	}

	flag.StringVar(&conf.HTTP.HostAddress, "a", conf.HTTP.HostAddress, "Endpoint server IP address host:port")
	flag.StringVar(&conf.MainDB.DSN, "d", sp.MainDB.DSN, "Database URI")
	flag.StringVar(&conf.AccurualService.Url, "r", conf.AccurualService.Url, "Accurual System Address")
	flag.StringVar(&conf.App.LogLevel, "l", conf.App.LogLevel, "Loging level")

	flag.Parse()

	if tStr, ok = os.LookupEnv("DATABASE_URI"); ok {
		fmt.Printf("LookupEnv(DATABASE_URI)=%v\n", tStr)
		conf.MainDB.DSN = tStr
	}

	fmt.Printf("Config parametrs:\n")
	fmt.Printf("Http server ADDRESS=%v\n", conf.HTTP.HostAddress)
	fmt.Printf("DATABASE_URI=%v\n", conf.MainDB.DSN)
	fmt.Printf("http AccurualSystemAddress=%v\n", conf.AccurualService.Url)
	fmt.Printf("LogLevel=%v\n", conf.App.LogLevel)
	fmt.Printf("AccurualTimeReset=%v\n", conf.AccurualService.TimeReset)
	fmt.Printf("IntervalAccurual=%v\n", conf.AccurualService.Interval)
	fmt.Printf("TimeoutContexDB=%v\n", conf.HTTP.TimeoutContexDB)

	return conf, nil
}
