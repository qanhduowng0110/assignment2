package routes

import (
	"entdemo/worker"
	"entdemo/api"
)

func Route() {
	api.ApiGet()
	worker.AutoSync()
}