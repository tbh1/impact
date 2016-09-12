package manage

import "github.com/gorilla/mux"

func Initialize(r *mux.Router) {
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/env", EnvHandler).Methods("GET")
	r.HandleFunc("/info", InfoHandler).Methods("GET")
}
