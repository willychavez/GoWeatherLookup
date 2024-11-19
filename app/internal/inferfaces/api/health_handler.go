package api

import (
	"fmt"
	"net/http"
	"time"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(time.Now())
	if duration.Seconds() > 10 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
