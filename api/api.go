package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tbh1/impact/api/manage"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func Initialize(r *mux.Router) {
	manage.Initialize(r.PathPrefix("/manage").Subrouter())
}
