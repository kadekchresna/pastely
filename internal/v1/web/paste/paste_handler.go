package paste

import (
	"net/http"

	"github.com/kadekchresna/pastely/internal/v1/usecase/paste"
	"github.com/labstack/echo/v4"
)

type pasteHandler struct {
	PasteUsecase paste.PasteUsecase
}

func NewPasteUsecase(PasteUsecase paste.PasteUsecase) PasteHandler {
	return &pasteHandler{
		PasteUsecase: PasteUsecase,
	}
}

func (h *pasteHandler) GetPaste(c echo.Context) error {
	ctx := c.Request().Context()

	var params paste.GetPasteParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	logs, err := h.PasteUsecase.GetPaste(ctx, params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": logs})
}
