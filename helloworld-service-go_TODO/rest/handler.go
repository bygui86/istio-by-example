package rest

import (
	"net/http"

	"github.com/bygui86/go-metrics/utils/logger"
	"github.com/gorilla/mux"
)

const (
	msgDefault = "Hello world!"
)

func echo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	msg := vars["msg"]

	if len(msg) == 0 {
		logger.Log.Infof("[REST] Echo of default msg '%s'", msgDefault)
		w.Write([]byte(msgDefault))
	} else {
		logger.Log.Infof("[REST] Echo of msg '%s'", msg)
		w.Write([]byte(msg))
	}
}

func echoWithMetrics(w http.ResponseWriter, r *http.Request, customMetrics *CustomMetrics) {

	vars := mux.Vars(r)
	msg := vars["msg"]

	if len(msg) == 0 {
		logger.Log.Infof("[REST] Echo of default msg '%s'", msgDefault)
		w.Write([]byte(msgDefault))
		go customMetrics.incrementOpsProcessed("default")
		go customMetrics.incrementEchoRequests()
	} else {
		logger.Log.Infof("[REST] Echo of msg '%s'", msg)
		w.Write([]byte(msg))
		go customMetrics.incrementOpsProcessed("msg")
		go customMetrics.incrementEchoMsgRequests()
	}
}
