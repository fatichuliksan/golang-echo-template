package main

import (
	"project/src/api"
	"project/src/helper"

	"github.com/labstack/echo"
)

func main() {
	api := api.NewAPI{
		Echo:   echo.New(),
		Helper: helper.NewHelper(),
	}
	api.Register()
	api.Echo.Start(api.Helper.Config.GetString(`app.host`))
}
