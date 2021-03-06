package api

import (
	"errors"
	"fmt"
)

type ErrorOut struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *ErrorOut) Error() string {
	return fmt.Sprintf("(%s) %s", e.Code, e.Message)
}

func GetError(err error) *ErrorOut {
	var e *ErrorOut
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func IsError(err error, code string) bool {
	apie := GetError(err)
	return apie != nil && apie.Code == code
}
