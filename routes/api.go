package routes

import (
	"youtube-manager-go/middlewares"
	"youtube-manager-go/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")

	{
		g.GET("/popular", api.FetchMostPupularVideos())
		g.GET("/video/:id", api.GetVideo(), middlewares.FirebaseAuth())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}

	fg := g.Group("/favorite", middlewares.FirebaseGuard())

	{
		fg.GET("", api.FetchFavoriteVideos())
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
	}
}