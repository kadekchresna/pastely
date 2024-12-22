package paste

import (
	"net/http"

	"github.com/kadekchresna/pastely/internal/v2/usecase/paste"
	"github.com/labstack/echo/v4"
)

type pasteHandler struct {
	PasteUsecase paste.PasteUsecase
}

func NewPasteHandler(PasteUsecase paste.PasteUsecase) PasteHandler {
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
	data, err := h.PasteUsecase.GetPaste(ctx, params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": data})
}

func (h *pasteHandler) CreatePaste(c echo.Context) error {
	ctx := c.Request().Context()

	var data paste.CreatePaste
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	p, err := h.PasteUsecase.CreatePaste(ctx, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": p})
}

func (h *pasteHandler) GetPresignedURL(c echo.Context) error {
	ctx := c.Request().Context()

	objectKey := c.QueryParam("object_key")
	if len(objectKey) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "objectKey is required"})
	}

	data, err := h.PasteUsecase.GetPresignedURL(ctx, objectKey, 60)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": data})
}

func (h *pasteHandler) PutPresignedURL(c echo.Context) error {
	ctx := c.Request().Context()

	objectKey := c.QueryParam("object_key")
	if len(objectKey) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "objectKey is required"})
	}

	data, err := h.PasteUsecase.PutPresignedURL(ctx, objectKey, 60)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": data})
}
