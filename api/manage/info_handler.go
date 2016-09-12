package manage

import (
	"net/http"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"syscall"
)

type InfoState struct {
	CpuUsage syscall.Rusage
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	usage := syscall.Rusage{}
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &usage); err != nil {
		log.WithError(err).Error("Failed to get usage")
	}

	body, err := json.Marshal(InfoState{
		CpuUsage: usage,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"request": r,
			"error": err,
		}).Error("Failed to marshal JSON response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}

