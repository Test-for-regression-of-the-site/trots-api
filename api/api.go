// Package api provides the trots API controller.
package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Test-for-regression-of-the-site/trots-api/pkg/logz"

	"github.com/go-chi/chi"
)

// API is the trots API controller.
type API struct {
	cfg    Config
	routes *chi.Mux
}

// NewAPI creates a new API controller or returns an error.
// It uses some reasonable defaults (serving address, etc).
// Options can configure the settings structure.
func NewAPI(opts ...Option) (*API, error) {
	var cfg = Config{
		Addr: "localhost:2020",
		Logger: func(name string) Logger {
			return logz.NewDev(os.Stderr).Named(name)
		},
	}
	for _, setOpt := range opts {
		var err = setOpt(&cfg)
		if err != nil {
			return nil, fmt.Errorf("setting option: %w", err)
		}
	}
	var api = &API{
		cfg:    cfg,
		routes: chi.NewMux(),
	}
	api.initRoutes()
	return api, nil
}

// ServeHTTP allows to use the API a simple http handler.
func (api *API) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	api.routes.ServeHTTP(rw, req)
}

// ListenAndServe starts a server at configured port.
func (api *API) ListenAndServe(ctx context.Context) error {
	var server = &http.Server{
		Addr:        api.cfg.Addr,
		Handler:     api,
		BaseContext: func(net.Listener) context.Context { return ctx },
		ErrorLog:    stdErrorLog(api.cfg.Logger("server")),

		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second / 2,
		WriteTimeout:      10 * time.Second,
	}
	return server.ListenAndServe()
}
