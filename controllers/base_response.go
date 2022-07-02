package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

type ErrResponse struct {
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

type OkResponse struct {
	Message interface{} `json:"message"`
}

func SuccessResponse(c echo.Context, data interface{}) error {
	response := DataResponse{}
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func SuccessOkResponse(c echo.Context) error {
	response := OkResponse{}
	response.Message = "OK"
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, status int, message string, errs error) error {
	response := ErrResponse{}
	response.Message = errs.Error()
	return c.JSON(http.StatusInternalServerError, response)
}
