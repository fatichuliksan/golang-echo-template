package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type responseHelper struct {
}

type responseFormat struct {
	C       echo.Context
	Code    int
	Status  string
	Message string
	Data    interface{}
}

// Method ...
type Method interface {
	SetResponse(c echo.Context, code int, status string, message string, data interface{}) responseFormat
	SendResponse(res responseFormat) error
	EmptyJSONMap() map[string]interface{}
	SendSuccess(c echo.Context, message string, data interface{}) error
	SendBadRequest(c echo.Context, message string, data interface{}) error
	SendError(c echo.Context, message string, data interface{}) error
	SendUnauthorized(c echo.Context, message string, data interface{}) error
	SendValidationError(c echo.Context, validationErrors validator.ValidationErrors) error
	SendCustomResponse(c echo.Context, httpCode int, status string, message string, data interface{}) error
}

// NewResponse ...
func NewResponse() Method {
	return &responseHelper{}
}

// SetResponse ...
func (r *responseHelper) SetResponse(c echo.Context, code int, status string, message string, data interface{}) responseFormat {
	return responseFormat{c, code, status, message, data}
}

// SendResponse ...
func (r *responseHelper) SendResponse(res responseFormat) error {
	if len(res.Message) == 0 {
		res.Message = http.StatusText(res.Code)
	}

	if res.Data != nil {
		return res.C.JSON(res.Code, map[string]interface{}{
			"code":    res.Code,
			"status":  res.Status,
			"message": res.Message,
			"data":    res.Data,
		})
	} else {
		return res.C.JSON(res.Code, map[string]interface{}{
			"code":    res.Code,
			"status":  res.Status,
			"message": res.Message,
		})
	}
}

// EmptyJSONMap : set empty data.
func (r *responseHelper) EmptyJSONMap() map[string]interface{} {
	return make(map[string]interface{})
}

// SendSuccess : Send success response to consumers.
func (r *responseHelper) SendSuccess(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusOK, "success", message, data)
	return r.SendResponse(res)
}

// SendBadRequest : Send bad request response to consumers.
func (r *responseHelper) SendBadRequest(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusBadRequest, "error", message, data)
	return r.SendResponse(res)
}

// SendError : Send error request response to consumers.
func (r *responseHelper) SendError(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusInternalServerError, "error", message, data)
	return r.SendResponse(res)
}

// SendUnauthorized : Send error request response to consumers.
func (r *responseHelper) SendUnauthorized(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusUnauthorized, "error", message, data)
	return r.SendResponse(res)
}

// SendValidationError : Send validation error request response to consumers.
func (r *responseHelper) SendValidationError(c echo.Context, validationErrors validator.ValidationErrors) error {
	errorResponse := []string{}
	for _, err := range validationErrors {
		errorResponse = append(errorResponse, strings.Trim(fmt.Sprint(err), "[]")+".")
	}
	res := r.SetResponse(c, http.StatusBadRequest, "error", strings.Trim(fmt.Sprint(errorResponse), "[]"), r.EmptyJSONMap())
	return r.SendResponse(res)
}

// SendCustomResponse ...
func (r *responseHelper) SendCustomResponse(c echo.Context, httpCode int, status string, message string, data interface{}) error {
	res := r.SetResponse(c, httpCode, status, message, data)
	return r.SendResponse(res)
}
