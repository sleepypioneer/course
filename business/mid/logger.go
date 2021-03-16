package mid

import (
	"context"
	"log"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func Logger(log *log.Logger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			log.Println("Started")

			if err := handler(ctx, w, r); err != nil {
				return err
			}

			log.Println("Completed")

			return nil
		}

		return h
	}

	return m
}
