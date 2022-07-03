package err

import (
	"errors"
	"net/http"
)

func ErrorGetUsersCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}

func ErrorAddUsersCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound
	} else if errors.Is(thisError, ErrNameEmpty) {
		return http.StatusBadRequest
	} else if errors.Is(thisError, ErrEmailEmpty) {
		return http.StatusBadRequest
	} else if errors.Is(thisError, ErrIDEmpty) {
		return http.StatusBadRequest
	} else if errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func ErrorUpdateUsersCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusNotFound
	} else if errors.Is(thisError, ErrEmailEmpty) {
		return http.StatusBadRequest
	} else if errors.Is(thisError, ErrIDEmpty) {
		return http.StatusBadRequest
	} else if errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest
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
