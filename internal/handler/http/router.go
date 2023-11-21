package http

import (
	"net/http/pprof"

	"github.com/go-chi/chi"
	midlw "github.com/go-chi/chi/middleware"
	"github.com/greyfox12/GoDilpom1/internal/config"
	"github.com/greyfox12/GoDilpom1/pkg/logger"
	"github.com/greyfox12/GoDilpom1/pkg/middleware"
)

type HandlerRouter interface {
	AddRoutes(r *chi.Mux)
	GetVersion() string
	GetContentType() string
}

type Router struct {
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{router: chi.NewRouter()}
}

func (r *Router) WithMetrics() *Router {
	//r.router.Use(promlib.NewMiddleware(promlib.DefHTTPRequestDurBuckets).Handler)
	//r.router.Use(tracing.NewHTTPMiddleware(opentracing.GlobalTracer()).Handler)

	return r
}

func (r *Router) WithHandler(h HandlerRouter, logger logger.Logger, cfg config.Config) *Router {

	//ct := h.GetContentType()
	//if h.GetContentType() != "" {
	//    api = api.Headers("Content-Type", "application/json; charset=UTF-8")
	//}

	//	api.Use(middleware.AddContextMiddleware(logger, cfg))
	r.router.Use(middleware.AccessLogMiddleware(logger))
	r.router.Use(midlw.StripSlashes)

	h.AddRoutes(r.router)

	return r
}

func (r *Router) WithProfiler() *Router {
	r.router.HandleFunc("/debug/pprof/", pprof.Index)
	// Not securely - so disable it
	// r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	r.router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.router.Handle("/debug/pprof/block", pprof.Handler("block"))
	r.router.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	r.router.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))

	return r
}
