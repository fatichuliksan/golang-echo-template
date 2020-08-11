package handler

import (
	"project/src/delivery/api/request"
	"project/src/helper"

	"github.com/labstack/echo"
)

// ExampleHandler ...
type ExampleHandler struct {
	Helper helper.Helper
}

// GetExample ...
func (t *ExampleHandler) GetExample(c echo.Context) error {
	return t.Helper.Response.SendSuccess(c, "Example Handler", nil)
}

// PostExample ...
func (t *ExampleHandler) PostExample(c echo.Context) error {
	var req request.PostExample

	if err := c.Bind(&req); err != nil {
		return t.Helper.Response.SendBadRequest(c, err.Error(), nil)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	return t.Helper.Response.SendSuccess(c, "Example Handler", req)
}
