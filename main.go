package main

import (
	"entdemo/api"
	"entdemo/worker"
)

func main() {
	api.ApiGet()
	go worker.AutoSync()
}
