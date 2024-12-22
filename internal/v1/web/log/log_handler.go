package log

import (
	"net/http"

	"github.com/kadekchresna/pastely/internal/v1/usecase/log"
	"github.com/labstack/echo/v4"
)

type logHandler struct {
	LogUsecase log.LogUsecase
}

func NewLogHandler(LogUsecase log.LogUsecase) LogHandler {
	return &logHandler{
		LogUsecase: LogUsecase,
	}
}

func (h *logHandler) GetLog(c echo.Context) error {
	ctx := c.Request().Context()

	var params log.GetLogParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	data, err := h.LogUsecase.GetLog(ctx, params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": data})
}
