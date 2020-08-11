package route

import (
	appMiddleware "project/src/delivery/api/middleware"
	"project/src/helper"
	"project/src/repository"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// NewRoute Handler
type NewRoute struct {
	Echo     *echo.Echo
	Helper   helper.Helper
	DBMaster *gorm.DB
	DBWms    *gorm.DB
	DBOms    *gorm.DB
}

// Register ...
func (t *NewRoute) Register() {
	groupV1 := t.Echo.Group("api/v1")
	// contoh group baru
	exampleGroup := groupV1.Group("/example")
	t.ExampleRoute(exampleGroup)

	groupV1.Use(appMiddleware.JWTWithConfig(appMiddleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(t.Helper.Config.GetString("jwt.secret")),
		UserRepo:      repository.NewUserRepo(t.DBWms),
	}))

	t.AuthRoute(groupV1)
}
