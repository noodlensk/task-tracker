package httperr

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/noodlensk/task-tracker/internal/common/errors"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorized(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Unauthorized", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Bad request", http.StatusBadRequest)
}

func NotFound(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Not found", http.StatusNotFound)
}

func RespondWithSlugError(err error, w http.ResponseWriter, r *http.Request) {
	slugError, ok := err.(errors.SlugError) //nolint:errorlint
	if !ok {
		InternalError("internal-server-error", err, w, r)

		return
	}

	switch slugError.ErrorType() {
	case errors.ErrorTypeAuthorization:
		Unauthorized(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeNotFound:
		NotFound(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}

//nolint:unparam
func httpRespondWithError(err error, slug string, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	resp := ErrorResponse{slug, status}

	if err := render.Render(w, r, resp); err != nil { //nolint:staticcheck
		// TODO: handle err
	}
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)

	return nil
}
