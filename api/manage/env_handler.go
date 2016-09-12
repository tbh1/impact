package manage

import (
	"net/http"
	"github.com/spf13/viper"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

func EnvHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: filter sensitive info
	body, err := json.Marshal(viper.AllSettings())
	if err != nil {
		log.WithFields(log.Fields{
			"request": r,
			"error": err,
		}).Error("Failed to marshal JSON response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}
