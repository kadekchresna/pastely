package web

import (
	"github.com/kadekchresna/pastely/internal/v1/web/log"
	"github.com/kadekchresna/pastely/internal/v1/web/paste"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Paste paste.PasteHandler
	Log   log.LogHandler
}

func InitAPI(
	e *echo.Echo,
	h Handlers,

) {
	v1 := e.Group("/api/v1")

	v1.GET("/paste", h.Paste.GetPaste)
	v1.POST("/paste", h.Paste.CreatePaste)
	v1.DELETE("/paste", h.Paste.DeleteExpiredPastes)

	v1.GET("/log", h.Log.GetLog)

}
