package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, comp templ.Component) error {
	return comp.Render(ctx.Request().Context(), ctx.Response())
}
