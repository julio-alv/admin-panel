package handler

import (
	"admin-panel/view/home"

	"github.com/labstack/echo/v4"
)

func HomeHandler(ctx echo.Context) error {
	return render(ctx, home.Show())
}
