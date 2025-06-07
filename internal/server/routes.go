package server

import (
	"context"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return mux
}

func StartServer(lc fx.Lifecycle, mux *http.ServeMux) {
	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Server failed: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
