package services

import (
	"context"
	"encoding/json"
	"entdemo/ent"
	"entdemo/model"
	"fmt"
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
		fmt.Println("Error:", err)
		return
	}

	//host=localhost port=5432 user=root dbname=postgres password=secret sslmode=disable sslmode=disable
	client, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=root dbname=postgres password=secret sslmode=disable sslmode=disable") //search_path=simple_bank
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	for i := 0; i < len(earthquake.Features); i++ {

		earthquake_report := client.Earthquake.Create().
			SetFeatureID(earthquake.Features[i].ID).
			SetMagnitude(earthquake.Features[i].Properties.Mag).
			SetOccurTime(convertIntToTimeStamp(earthquake.Features[i].Properties.Time)).
			SetUpdateTime(convertIntToTimeStamp(earthquake.Features[i].Properties.Updated)).
			SetURL(earthquake.Features[i].Properties.URL).
			SetDetailURL(earthquake.Features[i].Properties.Detail).
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
			SetEarthquakeType(earthquake.Features[i].Type).
			SaveX(context.Background())

		earthquake_report_id := earthquake_report.ID

		client.FeatureType.Create().
			SetEarthquakeID(earthquake_report_id).
			SetFeatureProductType(earthquake.Features[i].Type).
			SaveX(context.Background())

		client.Geometry.Create().
			SetEarthquakeID(earthquake_report_id).
			SetLongitude(earthquake.Features[i].Geometry.Coordinates[0]).
			SetLatitude(earthquake.Features[i].Geometry.Coordinates[1]).
			SetDepth(earthquake.Features[i].Geometry.Coordinates[2]).
			SetPlace(earthquake.Features[i].Properties.Place).
			SaveX(context.Background())

		client.AssociatedEvent.Create().
			SetEarthquakeID(earthquake_report_id).
			SetAssociateID(earthquake.Features[i].Properties.Ids).
			SaveX(context.Background())

		client.EventType.Create().
			SetEarthquakeID(earthquake_report_id).
			SetTypes(earthquake.Features[i].Geometry.Type).
			SaveX(context.Background())

		client.FeltReport.Create().
			SetEarthquakeID(earthquake_report_id).
			SetFelt(earthquake.Features[i].Properties.Felt).
			SetCdi(earthquake.Features[i].Properties.Cdi).
			SetMmi(earthquake.Features[i].Properties.Mmi).
			SetAlert(earthquake.Features[i].Properties.Alert).
			SaveX(context.Background())
	}
}
