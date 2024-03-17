package worker

function AutoSync(){
	app.Post("/api/earthquakes/sync-data", func(c *fiber.Ctx) error {
		services.ApiImportEarthquake()
		return c.SendString("Sucess!")
	})
	time.Sleep(10 * time.Minute)
}