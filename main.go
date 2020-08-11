package main

import (
	"flag"
	"log"
	"project/src/delivery/api"
	"project/src/helper"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	api := api.NewAPI{
		Echo:   echo.New(),
		Helper: helper.NewHelper(),
	}
	api.Register()

	port := flag.Int("port", api.Helper.Config.GetInt(`app.port`), "port")
	host := flag.String("host", api.Helper.Config.GetString(`app.host`), "host")
	flag.Parse()
	if api.Helper.Config.GetBool(`app.debug`) {
		log.Println("Service RUN on DEBUG mode - HOST: " + *host + ":" + strconv.Itoa(*port))
	}
	api.Echo.Start(*host + ":" + strconv.Itoa(*port))
}
