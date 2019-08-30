package rest

import (
	"github.com/bygui86/go-metrics/utils"
	"github.com/bygui86/go-metrics/utils/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	// Environment variables -
	restHostEnvVar            = "ECHOSERVER_REST_HOST"
	restPortEnvVar            = "ECHOSERVER_REST_PORT"
	restShutdownTimeoutEnvVar = "ECHOSERVER_REST_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	restHostDefault            = "localhost"
	restPortDefault            = 7001
	restShutdownTimeoutDefault = 15

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
	CustomMetrics   *CustomMetrics
}

// CustomMetrics - Container object for custom metrics
type CustomMetrics struct {
	opsProcessed    *prometheus.CounterVec
	echoRequests    prometheus.Counter
	echoMsgRequests prometheus.Counter
}

// newConfig -
func newConfig() (*Config, error) {

	logger.Log.Debugln("[REST] Setup new REST server config...")

	return &Config{
		RestHost:        utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort:        utils.GetIntEnv(restPortEnvVar, restPortDefault),
		ShutdownTimeout: utils.GetIntEnv(restShutdownTimeoutEnvVar, restShutdownTimeoutDefault),
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
		echoRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsGeneralNamespace,
			Subsystem: metricsRestSubsystem,
			Name:      echoRequestsMetricName,
			Help:      echoRequestsMetricHelp,
		}),
		echoMsgRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: metricsGeneralNamespace,
			Subsystem: metricsRestSubsystem,
			Name:      echoMsgRequestsMetricName,
			Help:      echoMsgRequestsMetricHelp,
		}),
	}
}

// IncrementOpsProcessed - Increment total number of processed operations
func (cm *CustomMetrics) incrementOpsProcessed(opType string) {

	cm.opsProcessed.WithLabelValues(opType).Inc()
}

// incrementEchoRequests - Increment total number of default echo requests
func (cm *CustomMetrics) incrementEchoRequests() {

	cm.echoRequests.Inc()
}

// incrementEchoMsgRequests - Increment total number of echo requests with message
func (cm *CustomMetrics) incrementEchoMsgRequests() {

	cm.echoMsgRequests.Inc()
}
