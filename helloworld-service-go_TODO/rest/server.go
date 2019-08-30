package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bygui86/go-metrics/utils/logger"

	"github.com/gorilla/mux"
)

// RestServer -
type RestServer struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
}

// NewRestServer - Create new REST server
func NewRestServer() (*RestServer, error) {

	logger.Log.Infoln("[REST] Setup new REST server...")

	// create config
	cfg, err := newConfig()
	if err != nil {
		return nil, err
	}

	// create router
	router := newRouter(cfg)

	// create http server
	httpServer := newHTTPServer(cfg.RestHost, cfg.RestPort, router)

	return &RestServer{
		Config:     cfg,
		Router:     router,
		HTTPServer: httpServer,
	}, nil
}

// newRouter -
func newRouter(cfg *Config) *mux.Router {

	logger.Log.Debugln("[REST] Setup new Router config...")

	router := mux.NewRouter().StrictSlash(false)
	// router.HandleFunc("/echo", echo).Methods(http.MethodGet)
	// router.HandleFunc("/echo/{msg}", echo).Methods(http.MethodGet)
	router.
		HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
			echoWithMetrics(w, r, cfg.CustomMetrics)
		}).
		Methods(http.MethodGet)
	router.
		HandleFunc("/echo/{msg}", func(w http.ResponseWriter, r *http.Request) {
			echoWithMetrics(w, r, cfg.CustomMetrics)
		}).
		Methods(http.MethodGet)
	return router
}

// newHttpServer -
func newHTTPServer(host string, port int, router *mux.Router) *http.Server {

	logger.Log.Debugf("[REST] Setup new HTTP server on port %d...", port)

	return &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

// Start - Start REST server
func (s *RestServer) Start() {

	logger.Log.Infoln("[REST] Start REST server...")

	// TODO add a channel to communicate if everything is right
	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil {
			logger.Log.Errorln("[REST] Error starting REST server:", err)
		}
	}()

	logger.Log.Infoln("[REST] REST server listen on port", s.Config.RestPort)
}

// Shutdown - Shutdown REST server
func (s *RestServer) Shutdown() {

	logger.Log.Warnln("[REST] Shutdown REST server...")
	if s.HTTPServer != nil {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Config.ShutdownTimeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		s.HTTPServer.Shutdown(ctx)
	}
}
