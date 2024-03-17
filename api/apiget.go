package api

import (
	"entdemo/ent"
	"entdemo/services"
	"reflect"
	"strconv"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
)

// convert interface to map[string] interface
func structToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	if objType.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i).Interface()
		result[field.Name] = fieldValue
	}

	return result
}

func ApiGet() {
	app := fiber.New()

	app.Get("/get-all", func(c *fiber.Ctx) error {
		var earthquakes = services.Service_get_db()
		return c.JSON(earthquakes)
	})

	app.Get("/get-by-paging", func(c *fiber.Ctx) error {
		var pageSize, errSize = strconv.Atoi(c.Query("pageSize"))
		var pageIndex, errIndex = strconv.Atoi(c.Query("pageIndex"))
		var earthquakes = services.Service_get_db_by_paging(pageIndex, pageSize)
		if errSize != nil || errIndex != nil {
			return c.SendStatus(400)
		}
		return c.JSON(earthquakes)
	})

	app.Post("/sync-data", func(c *fiber.Ctx) error {
		services.ApiImportEarthquake()
		return c.SendString("Sucess!")
	})
	app.Get("/filter-by-status", func(c *fiber.Ctx) error {
		var stt string = c.Query("status")
		var result []*ent.Earthquake = services.Service_filter_by_status(stt)
		c.SendString("Sucess!")
		return c.JSON(result)
	})

	cfg := swagger.Config{
		URL:          "localhost:3000",
		DeepLinking:  false,
		DocExpansion: "none",
		Title:        "Swagger API Docs",
	}

	app.Use(swagger.New(cfg))
	app.Listen(":3000")
}
