package rest

type Hello struct {
	Hostname string `json:"hostname"`
	Greeting string `json:"greeting"`
	Version  string `json:"version"`
}
