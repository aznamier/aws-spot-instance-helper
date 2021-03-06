package healthcheck

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	router          = mux.NewRouter()
	healthcheckPort = ":9777"
)

//StartHealthcheck - function to start health checking
func StartHealthcheck() {
	router.HandleFunc("/ping", healthCheck).Methods("GET", "HEAD").Name("Healthcheck")
	log.Info("Healthcheck handler is listening on ", healthcheckPort)
	log.Fatal(http.ListenAndServe(healthcheckPort, router))
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	// 1) test controller
	w.Write([]byte("pong"))
}
