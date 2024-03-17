package services

import (
	"context"
	"encoding/json"
	"entdemo/ent"
	"entdemo/model"
	"log"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	fetch "github.com/yngfoxx/gofiber-fetch"
)

func convertIntToTimeStamp(milliseconds int64) time.Time {
	baseTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	// Convert milliseconds to a duration
	duration := time.Duration(milliseconds) * time.Millisecond
	// Add the duration to the base time
	timestamp := baseTime.Add(duration)
	return timestamp

}

func ApiImportEarthquake() {

	url := "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_week.geojson"

	res := fetch.Method("GET").FiberFetch(url)
	if res.Error != nil {
		panic(res.Error)
	}
	var earthquake model.EarthquakeViewModel
	err := json.Unmarshal([]byte(res.Data.([]byte)), &earthquake)
	if err != nil {
		return
	}

	client, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=root dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	var ctx = context.Background()
	// Sync function
	// Delete all data in DB
	client.FtypeEarthquake.Delete().Exec(ctx)
	client.Earthquake.Delete().Exec(ctx)
	client.Geometry.Delete().Exec(ctx)
	client.Report.Delete().Exec(ctx)

	for i := 0; i < len(earthquake.Features); i++ {
		var reportNew = client.Report.Create().
			SetFelt(earthquake.Features[i].Properties.Felt).
			SetAlert(earthquake.Features[i].Properties.Alert).
			SetMmi(earthquake.Features[i].Properties.Mmi).
			SetCdi(earthquake.Features[i].Properties.Cdi).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			SaveX(ctx)

		var geoNew = client.Geometry.Create().
			SetLocationID(1).
			SetLongitude(earthquake.Features[i].Geometry.Coordinates[0]).
			SetLatitude(earthquake.Features[i].Geometry.Coordinates[1]).
			SetDepth(earthquake.Features[i].Geometry.Coordinates[2]).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			SaveX(ctx)
		var tzValue int32
		if tz, ok := earthquake.Features[i].Properties.Tz.(int32); ok {
			tzValue = tz
		}

		earthquake_report := client.Earthquake.Create().
			SetGeoID(geoNew.ID).
			SetReportID(reportNew.ID).
			SetMag(earthquake.Features[i].Properties.Mag).
			SetTime(convertIntToTimeStamp(earthquake.Features[i].Properties.Time)).
			SetUpdatedTime(convertIntToTimeStamp(earthquake.Features[i].Properties.Updated)).
			SetURL(earthquake.Features[i].Properties.URL).
			SetTz(int32(tzValue)).
			SetDetail(earthquake.Features[i].Properties.Detail).
			SetStatus(earthquake.Features[i].Properties.Status).
			SetTsunami(earthquake.Features[i].Properties.Tsunami).
			SetSig(earthquake.Features[i].Properties.Sig).
			SetNet(earthquake.Features[i].Properties.Net).
			SetCode(earthquake.Features[i].Properties.Code).
			SetNst(earthquake.Features[i].Properties.Nst).
			SetDmin(earthquake.Features[i].Properties.Dmin).
			SetRms(earthquake.Features[i].Properties.Rms).
			SetGap(earthquake.Features[i].Properties.Gap).
			SetMagType(earthquake.Features[i].Properties.MagType).
			SetEqType(earthquake.Features[i].Type).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			SaveX(ctx)

		earthquake_report_id := earthquake_report.ID

		client.FtypeEarthquake.Create().
			SetEarthquakeID(earthquake_report_id).
			SetFtID(1).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			SaveX(context.Background())
	}
}
