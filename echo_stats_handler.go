package echo_stats_handler

import (
	"encoding/json"
	"net/http"

	"github.com/fukata/golang-stats-api-handler"
	"github.com/labstack/echo"
)

func Stats(c echo.Context) error {
	bytes, err := json.Marshal(stats_api.GetStats())

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(bytes))
}
