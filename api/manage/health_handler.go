package manage

import (
	"net/http"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/tbh1/impact/db"
)

type HealthState struct {
	Healthy bool
	DatabaseHealth string
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	health := HealthState{Healthy: true}

	dbHealth, err := db.HealthCheck()
	if err != nil {
		health.Healthy = false
		health.DatabaseHealth = err.Error()
	}
	health.DatabaseHealth = dbHealth

	body, err := json.Marshal(health)
	if err != nil {
		log.WithFields(log.Fields{
			"request": r,
			"error": err,
		}).Error("Failed to marshal JSON response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}
