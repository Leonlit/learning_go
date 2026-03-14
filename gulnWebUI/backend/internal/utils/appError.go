package utils

import "log"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func BadRequest(code, msg string) *AppError {
	return NewAppError(
		code,
		msg,
		400,
		nil,
	)
}

func Unauthorized(msg string, err error) *AppError {
	return NewAppError(
		"UNAUTHORIZED",
		msg,
		401,
		err,
	)
}

func InternalError(msg string, err error) *AppError {
	return NewAppError(
		"INTERNAL_SERVER_ERROR",
		msg,
		500,
		err,
	)
}

func NewAppError(code string, message string, status int, err error) *AppError {
	appErr := &AppError{
		Code:    code,
		Message: message,
		Status:  status,
		Err:     err,
	}

	// Run logic here
	if err != nil {
		log.Println("AppError created:", code, message, status, err)
	}

	return appErr
}
