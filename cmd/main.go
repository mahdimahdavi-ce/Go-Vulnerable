package main

import (
	"ci/handlers"

	_ "github.com/Masterminds/semver/v3"
	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/prometheus/client_golang/prometheus"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.uber.org/zap"
	_ "google.golang.org/grpc"
)

func main() {

	app := fiber.New()

	// Loging all incoming requests
	app.Use(logger.New())

	app.Get("/", handlers.GetHandler)

	app.Listen(":3000")

}
