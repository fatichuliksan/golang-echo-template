package api

import (
	"fmt"
	"net/http"
	"project/src/delivery/api/route"
	"project/src/helper"
	"project/src/helper/postgre"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewAPI struct ...
type NewAPI struct {
	Echo   *echo.Echo
	Helper helper.Helper
}

// Register ...
func (t *NewAPI) Register() *NewAPI {
	t.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	t.Echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	t.Echo.HTTPErrorHandler = t.HTTPErrorHandlerCustom
	t.Echo.Validator = t.Helper.Validator.Validator()

	dbMaster := postgre.NewPostgre(t.Helper.Config.GetString("database.postgre.db_master.username"), t.Helper.Config.GetString("database.postgre.db_master.password"), t.Helper.Config.GetString("database.postgre.db_master.host"), t.Helper.Config.GetInt("database.postgre.db_master.port"), t.Helper.Config.GetString("database.postgre.db_master.database"))
	dbWms := postgre.NewPostgre(t.Helper.Config.GetString("database.postgre.db_wms.username"), t.Helper.Config.GetString("database.postgre.db_wms.password"), t.Helper.Config.GetString("database.postgre.db_wms.host"), t.Helper.Config.GetInt("database.postgre.db_wms.port"), t.Helper.Config.GetString("database.postgre.db_wms.database"))
	dbOms := postgre.NewPostgre(t.Helper.Config.GetString("database.postgre.db_oms.username"), t.Helper.Config.GetString("database.postgre.db_oms.password"), t.Helper.Config.GetString("database.postgre.db_oms.host"), t.Helper.Config.GetInt("database.postgre.db_oms.port"), t.Helper.Config.GetString("database.postgre.db_oms.database"))

	connDbMaster, _ := dbMaster.Connect()
	connDbWms, _ := dbWms.Connect()
	connDbOms, _ := dbOms.Connect()

	if true == t.Helper.Config.GetBool(`app.debug`) {
		// connDbMaster.LogMode(true)
		// connDbWms.LogMode(true)
		// connDbOms.LogMode(true)
		t.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "remote:${remote_ip}, method:${method}, uri:${uri}, status:${status}, latency:${latency_human}, error:${error}\n",
		}))
		// t.Echo.Use(middleware.Logger())
		t.Echo.HideBanner = true
		t.Echo.Debug = true
	} else {
		t.Echo.HideBanner = true
		t.Echo.Debug = false
		t.Echo.Use(middleware.Recover())
	}
	route := route.NewRoute{
		Echo:     t.Echo,
		Helper:   t.Helper,
		DBMaster: connDbMaster,
		DBWms:    connDbWms,
		DBOms:    connDbOms,
	}
	route.Register()

	return t
}

// HTTPErrorHandlerCustom ...
func (t *NewAPI) HTTPErrorHandlerCustom(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		report.Code = http.StatusBadRequest
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			}
			break
		}
	}
	t.Helper.Response.SendCustomResponse(c, report.Code, "error", report.Message.(string), nil)
}
