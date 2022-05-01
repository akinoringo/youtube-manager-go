package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func FetchMostPupularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(fasthttp.StatusOK, "Most Popular")
	}
}