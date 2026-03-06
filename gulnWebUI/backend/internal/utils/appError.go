package utils

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    string // machine-readable
	Message string // client-readable
	Status  int    // HTTP status code
}

func (e *AppError) Error() string {
	return e.Message
}

func WriteError(w http.ResponseWriter, err error) {
	var appErr *AppError

	if errors.As(err, &appErr) {
		SendJSONResponse(w, map[string]interface{}{
			"error":   appErr.Code,
			"message": appErr.Message,
		}, appErr.Status)
		return
	}

	// fallback: unexpected error
	SendJSONResponse(w, map[string]interface{}{
		"error":   "INTERNAL_SERVER_ERROR",
		"message": "Something went wrong",
	}, http.StatusInternalServerError)
}
