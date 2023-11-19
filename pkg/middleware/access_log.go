package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/greyfox12/GoDilpom1/pkg/logger"
	"go.uber.org/zap"
)

func AccessLogMiddleware(log logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//ctx := r.Context()
			start := time.Now()
			responseWriterWrapper := newResponseWriterWrapper(w)

			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Error("Failed to read request body")
			}

			if len(reqBody) > 0 {
				err = r.Body.Close() //  must close
				if err != nil {
					log.Error("Failed to close body reader")
				} else {
					r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
				}
			}

			defer func() {
				fields := make(map[string]interface{})
				//fields := make(map[string]string)
				pnc := recover()
				if pnc != nil {
					fields["event"] = "recovered after panic"
					//			fields["panic_value"] = pnc
					fields["stacktrace"] = string(debug.Stack())

					responseWriterWrapper.WriteHeader(http.StatusInternalServerError)
				}
				fields["duration"] = time.Since(start).String()
				fields["method"] = r.Method
				fields["path"] = r.URL.Path
				fields["status_code"] = responseWriterWrapper.StatusCode()

				//if responseWriterWrapper.statusCode == http.StatusBadRequest {
				fields["response_body"] = string(responseWriterWrapper.Body())
				fields["request_body"] = string(reqBody)
				//}

				if responseWriterWrapper.statusCode == http.StatusInternalServerError {
					log.Error("access_log")
				} else {
					log.Info("access_log", zap.String("duration", time.Since(start).String()),
						zap.String("method", r.Method),
						zap.String("path", r.URL.Path),
						zap.Int("status_code", responseWriterWrapper.StatusCode()),
						zap.String("response_body", string(responseWriterWrapper.Body())),
						zap.String("request_body", string(reqBody)))
				}
			}()

			next.ServeHTTP(responseWriterWrapper, r)
		})
	}
}
