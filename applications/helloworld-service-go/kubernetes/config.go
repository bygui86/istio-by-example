package kubernetes

import (
	"github.com/bygui86/helloworld-service-go/logger"
	"github.com/bygui86/helloworld-service-go/utils"
)

const (
	// Environment variables -
	kubeHostEnvVar            = "HELLOSVC_KUBE_HOST"
	kubePortEnvVar            = "HELLOSVC_KUBE_PORT"
	kubeShutdownTimeoutEnvVar = "HELLOSVC_KUBE_SHUTDOWN_TIMEOUT"

	// Default values -
	// host values: '0.0.0.0' for kubernetes, 'localhost' for local
	kubeHostDefault     = "localhost"
	kubePortDefault     = 8091
	kubeShutdownTimeout = 15
)

// Config -
type Config struct {
	RestHost        string
	RestPort        int
	ShutdownTimeout int
}

// newConfig -
func newConfig() (*Config, error) {

	logger.Log.Debugln("[KUBERNETES] Setup new Kubernetes config...")

	return &Config{
		RestHost:        utils.GetStringEnv(kubeHostEnvVar, kubeHostDefault),
		RestPort:        utils.GetIntEnv(kubePortEnvVar, kubePortDefault),
		ShutdownTimeout: utils.GetIntEnv(kubeShutdownTimeoutEnvVar, kubeShutdownTimeout),
	}, nil
}
