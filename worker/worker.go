package worker

import (
	"entdemo/services"
	"time"
)

func AutoSync() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			services.ApiImportEarthquake()
		}
	}
}
