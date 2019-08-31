package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bygui86/helloworld-service-go/logger"

	"github.com/gorilla/mux"
)

const (
	msgDefault = "Hello world!"
)

// hello -
func hello(w http.ResponseWriter, r *http.Request, cfg *Config) {

	vars := mux.Vars(r)
	name := vars["name"]

	hostname := getHostname()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		Hello{
			Hostname: hostname,
			Greeting: fmt.Sprintf(cfg.Greeting, name, hostname, cfg.Version),
			Version:  cfg.Version,
		},
	)

	logger.Log.Infof("[REST] Hello %s", name)

	incrementMetrics(cfg.CustomMetrics)
}

// getHostname -
func getHostname() string {

	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hostname
}

// incrementMetrics -
func incrementMetrics(customMetrics *CustomMetrics) {

	go customMetrics.incrementOpsProcessed("name")
	go customMetrics.incrementHelloRequests()
}
