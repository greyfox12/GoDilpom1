package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
// Структуры для обработки логирования HTTP запросов
type (

	// берём структуру для хранения сведений об ответе
	responseData struct {
		status int
		size   int
	}

	// добавляем реализацию http.ResponseWriter
	loggingResponseWriter struct {
		http.ResponseWriter // встраиваем оригинальный http.ResponseWriter
		responseData        *responseData
	}

)

	func (r *loggingResponseWriter) Write(b []byte) (int, error) {
		// записываем ответ, используя оригинальный http.ResponseWriter
		size, err := r.ResponseWriter.Write(b)
		r.responseData.size += size // захватываем размер
		return size, err
	}

	func (r *loggingResponseWriter) WriteHeader(statusCode int) {
		// записываем код статуса, используя оригинальный http.ResponseWriter
		r.ResponseWriter.WriteHeader(statusCode)
		r.responseData.status = statusCode // захватываем код статуса
	}
*/
type Logger interface {
	Debug(string, ...zapcore.Field)
	Info(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
	Fatal(string, ...zapcore.Field)
	Warn(string, ...zapcore.Field)
}

//type Log struct {
//loger *zap.Logger
//}

func NewLogger(logLevel string) (*zap.Logger, error) {
	//	l := zap.NewNop()
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		return nil, err
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout", "./logs.txt"}
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	// устанавливаем уровень
	cfg.Level = lvl

	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	// устанавливаем синглтон

	return zl, nil

}

/* Перенесно в access_log b response_writter_wrapper
// RequestLogger — middleware-логер для входящих HTTP-запросов.
func (l *Log) RequestLogger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := loggingResponseWriter{
			ResponseWriter: w, // встраиваем оригинальный http.ResponseWriter
			responseData:   responseData,
		}

		h(&lw, r)

		duration := time.Since(start)

		l.loger.Info("got incoming HTTP request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Duration("duration", duration),
			zap.Int("size", responseData.size),
			zap.Int("status", responseData.status),
		)
		l.loger.Sync()
		//	fmt.Printf("Write log\n")

	})
}

*/
