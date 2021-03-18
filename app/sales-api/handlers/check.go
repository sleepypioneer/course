package handlers

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"os"

	"github.com/ardanlabs/service/business/validate"
	"github.com/ardanlabs/service/foundation/web"
)

type checkGroup struct {
	build string
}

func (checkGroup) test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return validate.NewRequestError(errors.New("trusted error"), http.StatusNotFound)
	}

	status := struct {
		Status string
	}{
		Status: "OK-2",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

func (checkGroup) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK-2",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

func (cg checkGroup) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	info := struct {
		Status    string `json:"status,omitempty"`
		Build     string `json:"build,omitempty"`
		Host      string `json:"host,omitempty"`
		Pod       string `json:"pod,omitempty"`
		PodIP     string `json:"podIP,omitempty"`
		Node      string `json:"node,omitempty"`
		Namespace string `json:"namespace,omitempty"`
	}{
		Status:    "up",
		Build:     cg.build,
		Host:      host,
		Pod:       os.Getenv("KUBERNETES_PODNAME"),
		PodIP:     os.Getenv("KUBERNETES_NAMESPACE_POD_IP"),
		Node:      os.Getenv("KUBERNETES_NODENAME"),
		Namespace: os.Getenv("KUBERNETES_NAMESPACE"),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}
