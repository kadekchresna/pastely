package log

import "github.com/labstack/echo/v4"

type LogHandler interface {
	GetLog(c echo.Context) error
}
