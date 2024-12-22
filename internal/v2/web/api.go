package web

import (
	"github.com/kadekchresna/pastely/internal/v2/web/paste"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Paste paste.PasteHandler
}

func InitAPI(
	e *echo.Echo,
	h Handlers,

) {
	v2 := e.Group("/api/v2")

	v2.GET("/paste", h.Paste.GetPaste)
	v2.GET("/paste/pre-object", h.Paste.GetPresignedURL)
	v2.POST("/paste/pre-object", h.Paste.PutPresignedURL)
	v2.POST("/paste", h.Paste.CreatePaste)

}
