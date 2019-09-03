package rest

import (
	"github.com/bygui86/helloworld-service-go/logger"
	"github.com/bygui86/helloworld-service-go/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	// Environment variables -
	restHostEnvVar            = "HELLOSVC_REST_HOST"
	restPortEnvVar            = "HELLOSVC_REST_PORT"
	restShutdownTimeoutEnvVar = "HELLOSVC_REST_SHUTDOWN_TIMEOUT"
	restGreeting              = "HELLOSVC_REST_GREETING"
	restVersion               = "HELLOSVC_REST_VERSION"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	restHostDefault            = "localhost"
	restPortDefault            = 8080
	restShutdownTimeoutDefault = 15
	restGreetingDefault        = "Hello %s from %s with %s"
	restVersionDefault         = "1.0"

	// Custom metrics -
	metricsGeneralNamespace   = "general"
	metricsRestSubsystem      = "rest"
	opsProcessedMetricName    = "echoserver_processed_ops_total"
	opsProcessedMetricHelp    = "Total number of processed operations"
	echoRequestsMetricName    = "echoserver_echo_requests_total"
	echoRequestsMetricHelp    = "Total number of default echo requests"
	echoMsgRequestsMetricName = "echoserver_echo_msg_requests_total"
	echoMsgRequestsMetricHelp = "Total number of echo requests with messages"
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
	Greeting        string
	Version         string
	CustomMetrics   *CustomMetrics
}

// CustomMetrics - Container object for custom metrics
type CustomMetrics struct {
	opsProcessed  *prometheus.CounterVec
	helloRequests prometheus.Counter
}

// newConfig -
func newConfig() (*Config, error) {

	logger.Log.Debugln("[REST] Setup new REST server config...")

	return &Config{
		RestHost:        utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort:        utils.GetIntEnv(restPortEnvVar, restPortDefault),
		ShutdownTimeout: utils.GetIntEnv(restShutdownTimeoutEnvVar, restShutdownTimeoutDefault),
		Greeting:        utils.GetStringEnv(restGreeting, restGreetingDefault),
		Version:         utils.GetStringEnv(restVersion, restVersionDefault),
		CustomMetrics:   newCustomMetrics(),
	}, nil
}

// NewCustomMetrics -
func newCustomMetrics() *CustomMetrics {

	return &CustomMetrics{
		opsProcessed: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: metricsGeneralNamespace,
				Subsystem: metricsRestSubsystem,
				Name:      opsProcessedMetricName,
				Help:      opsProcessedMetricHelp,
			},
			[]string{"type"},
		),
		helloRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsGeneralNamespace,
			Subsystem: metricsRestSubsystem,
			Name:      echoRequestsMetricName,
			Help:      echoRequestsMetricHelp,
		}),
	}
}

// IncrementOpsProcessed - Increment total number of processed operations
func (cm *CustomMetrics) incrementOpsProcessed(opType string) {

	cm.opsProcessed.WithLabelValues(opType).Inc()
}

// incrementNameRequests - Increment total number of echo requests with message
func (cm *CustomMetrics) incrementHelloRequests() {

	cm.helloRequests.Inc()
}
