package api

import (
	"fmt"
	"net/http"
)

func (api *API) initRoutes() {
	api.routes.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(rw, "hello, %q", req.RemoteAddr)
	})
}
