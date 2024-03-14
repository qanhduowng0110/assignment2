package main

import (
	"entdemo/model"
	"entdemo/services"
	"math"
	"strconv"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/earthquakes/get-all", func(c *fiber.Ctx) error {
		var earthquakes = services.Service_get_db()
		return c.JSON(earthquakes)
	})
	app.Get("/api/earthquakes/get-by-paging", func(c *fiber.Ctx) error {
		var pageSize, errSize = strconv.Atoi(c.Query("pageSize"))
		var pageIndex, errIndex = strconv.Atoi(c.Query("pageIndex"))
		var earthquakes = services.Service_get_db()
		if errSize != nil || errIndex != nil {
			return c.SendStatus(400)
		}
		// bien tinh vi tri index cua thang thu dau tien trong trang
		startIndex := (pageIndex - 1) * pageSize

		// bien tinh vi tri index cua thang thu cuoi cung trong trang
		endIndex := pageIndex * pageSize

		if startIndex < 0 {
			startIndex = 0
		}

		if endIndex > len(earthquakes) {
			endIndex = len(earthquakes)
		}
		// tinh lay ra so data can lay trong mang arr
		page := earthquakes[startIndex:endIndex]
		var total = len(earthquakes)

		response := model.ResponseGetModel{
			Page:      page,
			Total:     total,
			TotalPage: int(math.Ceil(float64(total) / float64(pageSize))),
		}
		return c.JSON(response)
	})

	app.Get("/api/earthquakes/get-by-paging-v2", func(c *fiber.Ctx) error {
		var pageSize, errSize = strconv.Atoi(c.Query("pageSize"))
		var pageIndex, errIndex = strconv.Atoi(c.Query("pageIndex"))
		var earthquakes = services.Service_get_db_by_paging(pageIndex, pageSize)
		if errSize != nil || errIndex != nil {
			return c.SendStatus(400)
		}
		return c.JSON(earthquakes)
	})

	app.Post("/api/earthquakes/get-by-paging-v3", func(c *fiber.Ctx) error {
		var model model.EarthquakeFilterModel

		if err := c.BodyParser(model); err != nil {
			return err
		}
		var earthquakes = services.Service_get_clause_db_by_paging(model)
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
	app.Listen(":3001")
}
