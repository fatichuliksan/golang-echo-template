package route

import (
	"project/src/delivery/api/handler"

	"github.com/labstack/echo"
)

// ExampleRoute ...
func (t *NewRoute) ExampleRoute(g *echo.Group) {
	h := handler.ExampleHandler{
		Helper: t.Helper,
	}
	g.GET("", h.GetExample)
	g.POST("", h.PostExample)
}
