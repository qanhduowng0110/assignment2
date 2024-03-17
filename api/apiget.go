package api

import (
	"entdemo/model"
	"entdemo/services"
	"strconv"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
)

func ApiGet() {
	// services.ApiImportEarthquake()
	
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
	
	// api thu test func filter  muon filter them truong gi thi viet them vao model  EarthquakeFilterModel
	// viet them 1 method POST de nhan param tu payload body hien api nay dang fix cung
	app.Get("/api/earthquakes/get-by-paging-v3", func(c *fiber.Ctx) error {
	
		var equalFilter model.EarthquakeFilterModel
	
		// da fix cung o day
		equalFilter.PageIndex = 1
		equalFilter.PageSize = 10
		equalFilter.UpdateTimeFrom = 1710611748300
		equalFilter.UpdateTimeTo = 1710611838484
		var logfilter model.ApiReqFilterModel
		logfilter.ReqBody = equalFilter
	
		// luu lai log data luu vao log api
		// sau khi luu xong viet them 1 api get ra bo loc co tim kiem nhieu nhat trong ngay
		services.InsertLogApiRequest(logfilter)
		var earthquakes = services.Service_get_clause_db_by_paging(equalFilter)
	
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