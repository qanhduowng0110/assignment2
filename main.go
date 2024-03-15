package main

import (
	"entdemo/services"
	"strconv"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	services.ApiImportEarthquake()
	app := fiber.New()

	app.Get("/api/earthquakes/get-all", func(c *fiber.Ctx) error {
		var earthquakes = services.Service_get_db()
		return c.JSON(earthquakes)
	})

	app.Get("/api/earthquakes/get-by-paging", func(c *fiber.Ctx) error {
		var pageSize, errSize = strconv.Atoi(c.Query("pageSize"))
		var pageIndex, errIndex = strconv.Atoi(c.Query("pageIndex"))
		var earthquakes = services.Service_get_db_by_paging(pageIndex, pageSize)
		if errSize != nil || errIndex != nil {
			return c.SendStatus(400)
		}
		return c.JSON(earthquakes)
	})

	cfg := swagger.Config{
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		Title:        "Swagger API Docs",
	}

	app.Use(swagger.New(cfg))
	app.Listen(":3000")
}
