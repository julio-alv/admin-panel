package main

import (
	"admin-panel/handler"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	e.GET("/", handler.HomeHandler, withSession)
	e.GET("/login", handler.LoginHandler)

	e.POST("/login", func(ctx echo.Context) error {
		cred := Credentials{}
		err := ctx.Bind(&cred)
		if err != nil {
			return ctx.JSON(400, map[string]string{
				"message": "Invalid Request",
			})
		}
		cookie := &http.Cookie{}
		cookie.Name = "session"
		cookie.Value = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGlhcyI6ImxrYWFzIiwiYXVkIjoidGVuYW50IiwiZW1haWwiOiJ0ZXN0QGVtYWlsLmNvbSIsImV4cCI6MTcxMDk2MzI1NCwiaWF0IjoxNzEwNTMxMjU0LCJpZF9tZW1iZXIiOiJtZW1iZXJfVU1pTzY4b2pCZUFGMjZnaCIsImlkX3RlbmFudCI6InRlbmFudF9YVERrRmNhQ2xVaTY2T2dpIiwiaXNzIjoibG9jYWxob3N0OjgwODAiLCJzdWIiOiJtZW1iZXJfVU1pTzY4b2pCZUFGMjZnaCIsInR5cGUiOiJhY2Nlc3MifQ.l8ZIG0iCTuTXVyx6vMPghipgWkF7u57-cTgjfdBZy-0"
		cookie.Expires = time.Now().Add(2 * time.Hour)
		cookie.HttpOnly = true
		cookie.Secure = true
		cookie.SameSite = http.SameSiteLaxMode
		ctx.SetCookie(cookie)
		return ctx.Redirect(302, "/")
	})

	e.POST("/logout", func(ctx echo.Context) error {
		cookie := &http.Cookie{}
		cookie.Name = "session"
		cookie.Value = ""
		cookie.Expires = time.Now().Add(-time.Hour)
		ctx.SetCookie(cookie)
		return ctx.Redirect(302, "/login")
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func withSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("session")
		if err != nil {
			return ctx.Redirect(302, "/login")
		}
		if cookie.Value == "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbGlhcyI6ImxrYWFzIiwiYXVkIjoidGVuYW50IiwiZW1haWwiOiJ0ZXN0QGVtYWlsLmNvbSIsImV4cCI6MTcxMDk2MzI1NCwiaWF0IjoxNzEwNTMxMjU0LCJpZF9tZW1iZXIiOiJtZW1iZXJfVU1pTzY4b2pCZUFGMjZnaCIsImlkX3RlbmFudCI6InRlbmFudF9YVERrRmNhQ2xVaTY2T2dpIiwiaXNzIjoibG9jYWxob3N0OjgwODAiLCJzdWIiOiJtZW1iZXJfVU1pTzY4b2pCZUFGMjZnaCIsInR5cGUiOiJhY2Nlc3MifQ.l8ZIG0iCTuTXVyx6vMPghipgWkF7u57-cTgjfdBZy-0" {
			return next(ctx)
		}
		return ctx.Redirect(302, "/login")
	}
}
