// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ardanlabs/service/business/auth"
	"github.com/ardanlabs/service/business/mid"
	"github.com/ardanlabs/service/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	cg := checkGroup{build: build}
	app.Handle(http.MethodGet, "/test", cg.test, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodGet, "/debug/readiness", cg.readiness)
	app.Handle(http.MethodGet, "/debug/liveness", cg.liveness)

	return app
}
