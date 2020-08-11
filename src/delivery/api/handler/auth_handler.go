package handler

import (
	"project/src/delivery/api/request"
	"project/src/helper"
	"project/src/model"
	"project/src/usecase"

	"github.com/labstack/echo"
)

// AuthHandler ...
type AuthHandler struct {
	Helper      helper.Helper
	AuthUsecase usecase.AuthUsecase
}

// PostLogin ...
func (t *AuthHandler) PostLogin(c echo.Context) error {
	var (
		err error
		req request.PostLogin
	)

	if err = c.Bind(&req); err != nil {
		return t.Helper.Response.SendBadRequest(c, err.Error(), nil)
	}

	if err = c.Validate(req); err != nil {
		return err
	}

	token, err := t.AuthUsecase.Login(req.Email, req.Password)
	if err != nil {
		return t.Helper.Response.SendBadRequest(c, err.Error(), nil)
	}
	return t.Helper.Response.SendSuccess(c, "", map[string]interface{}{
		"token": token,
	})
}

// PostRefresh ...
func (t *AuthHandler) PostRefresh(c echo.Context) error {
	data := c.Get("user").(model.User)
	token, err := t.AuthUsecase.Refresh(data.ID)
	if err != nil {
		return t.Helper.Response.SendBadRequest(c, err.Error(), nil)
	}
	return t.Helper.Response.SendSuccess(c, "", map[string]interface{}{
		"token": token,
	})
}
