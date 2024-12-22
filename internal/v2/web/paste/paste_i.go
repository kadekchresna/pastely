package paste

import "github.com/labstack/echo/v4"

type PasteHandler interface {
	GetPaste(c echo.Context) error
	CreatePaste(c echo.Context) error
	GetPresignedURL(c echo.Context) error
	PutPresignedURL(c echo.Context) error
}
