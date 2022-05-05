package main

import (
	"youtube-manager-go/middlewares"
	"youtube-manager-go/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middlewares.YoutubeService())
	e.Use(middlewares.DatabaseService())
	e.Use(middlewares.Firebase())

	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}