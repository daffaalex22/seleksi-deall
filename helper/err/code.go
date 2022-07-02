package err

import (
	"errors"
	"net/http"
)

func ErrorGetUsersCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorAddUsersCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorUpdateUsersCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrDeleteUsersCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrUnathorizedCheck(thisError error) int {
	if errors.Is(thisError, ErrUnauthorized) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}
