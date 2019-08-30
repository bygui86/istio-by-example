package monitoring

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bygui86/go-metrics/utils/logger"

	"github.com/gorilla/mux"
)

// MonitorServer -
type MonitorServer struct {
	Config     *Config
	Router     *mux.Router
	HTTPServer *http.Server
}

// NewMonitorServer - Create new Monitoring REST server
func NewMonitorServer() (*MonitorServer, error) {

	logger.Log.Infoln("[MONITORING] Setup new REST server...")

	// create config
	cfg, err := newConfig()
	if err != nil {
		return nil, err
	}

	// create router
	router := newRouter()

	// create http server
	httpServer := newHTTPServer(cfg.RestHost, cfg.RestPort, router)

	return &MonitorServer{
		Config:     cfg,
		Router:     router,
		HTTPServer: httpServer,
	}, nil
}

// newRouter -
func newRouter() *mux.Router {

	logger.Log.Debugln("[MONITORING] Setup new Router config...")

	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/metrics", getMetricsHandler())
	return router
}

// newHttpServer -
func newHTTPServer(host string, port int, router *mux.Router) *http.Server {

	logger.Log.Debugf("[MONITORING] Setup new HTTP server on port %d...", port)

	return &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}

// Start - Start Monitoring REST server
func (s *MonitorServer) Start() {

	logger.Log.Infoln("[MONITORING] Start REST server...")

	// TODO add a channel to communicate if everything is right
	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil {
			logger.Log.Errorln("[MONITORING] Error starting REST server:", err)
		}
	}()

	logger.Log.Infoln("[MONITORING] REST server listen on port", s.Config.RestPort)
}

// Shutdown - Shutdown Monitoring REST server
func (s *MonitorServer) Shutdown() {

	logger.Log.Warnln("[MONITORING] Shutdown REST server...")
	if s.HTTPServer != nil {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.Config.ShutdownTimeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		s.HTTPServer.Shutdown(ctx)
	}
}
