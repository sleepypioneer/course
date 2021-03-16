package handlers

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK-2",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
