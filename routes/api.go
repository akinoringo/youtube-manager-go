package routes

import (
	"youtube-manager-go/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")

	{
		g.GET("/popular", api.FetchMostPupularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
	}
}