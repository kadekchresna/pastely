package paste

import "github.com/labstack/echo/v4"

type PasteHandler interface {
	GetPaste(c echo.Context) error
	CreatePaste(c echo.Context) error
	DeleteExpiredPastes(c echo.Context) error
}
