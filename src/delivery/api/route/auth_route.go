package route

import (
	"project/src/delivery/api/handler"
	"project/src/repository"
	"project/src/usecase"

	"github.com/labstack/echo"
)

// AuthRoute ...
func (t *NewRoute) AuthRoute(g *echo.Group) {
	// USECASE DEFINE
	authUsecase := usecase.AuthUsecase{
		Helper:   t.Helper,
		UserRepo: repository.NewUserRepo(t.DBWms),
	}

	h := handler.AuthHandler{
		Helper:      t.Helper,
		AuthUsecase: authUsecase,
	}
	g.POST("/login", h.PostLogin)
	g.POST("/refresh", h.PostRefresh)
}
