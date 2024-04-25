package handler

import (
	"admin-panel/view/login"

	"github.com/labstack/echo/v4"
)

func LoginHandler(ctx echo.Context) error {
	failed := ctx.QueryParam("failed")
	if failed == "true" {
		return render(ctx, login.Show(true))
	}
	return render(ctx, login.Show(false))
}
