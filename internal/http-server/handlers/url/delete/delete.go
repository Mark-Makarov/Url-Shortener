package delete

import (
	resp "Url-Shortener/internal/lib/api/response"
	"Url-Shortener/internal/lib/logger/sl"
	"Url-Shortener/internal/storage"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"path"
)

const aliasLength = 10

type Response struct {
	resp.Response
}

type URLDelete interface {
	DeleteURL(alias string) (string, int64, error)
}

func New(log *slog.Logger, urlDelete URLDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := path.Base(r.URL.Path)
		if alias == "" || len(alias) > aliasLength {
			log.Error("invalid alias")

			render.JSON(w, r, resp.Error("invalid alias"))

			return
		}

		_, id, err := urlDelete.DeleteURL(alias)
		if errors.Is(err, storage.ErrAliasNotFound) {
			log.Info("alias not found", slog.String("alias", alias))

			render.JSON(w, r, resp.Error("alias not found"))

			return
		}
		if err != nil {
			log.Error("alias not found", sl.Err(err))

			render.JSON(w, r, resp.Error("alias not found"))

			return
		}

		log.Info("alias deleted", slog.Int64("id", id))

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
