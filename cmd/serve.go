package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/gorilla/mux"
	"github.com/tbh1/impact/api"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/tbh1/impact/db"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Impact server",
	Long: `A longer description ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		fmt.Printf("args = %v", args)
		r := mux.NewRouter()

		// API base router
		api.Initialize(r.PathPrefix("/api").Subrouter())

		// Static content
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static")))).Name("static")

		srv := &http.Server{
			Handler:      logMiddleware(r),
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.WithFields(log.Fields{
			"address": srv.Addr,
		}).Info("Impact server is now running")

		// Set up database
		db.Initialize(db.Config{
			Type: db.MySQLType,
			Url: "impact:impact@tcp(127.0.0.1:3306)/impact",
		})

		log.Fatal(srv.ListenAndServe())
	},
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"remote": r.RemoteAddr,
			"request": r.RequestURI,
		}).Info()
		h.ServeHTTP(w, r)
	})
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Flags and configuration settings.
	serveCmd.PersistentFlags().IntP("port", "p", 8080, "Port to run on")
	serveCmd.PersistentFlags().BoolP("daemon", "d", false, "Run as daemon")
}
