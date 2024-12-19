package web

import (
	"github.com/kadekchresna/pastely/internal/v1/web/paste"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Paste paste.PasteHandler
}

func InitAPI(
	e *echo.Echo,
	h Handlers,

) {
	v1 := e.Group("/api/v1")

	v1.GET("/paste", h.Paste.GetPaste)

}
