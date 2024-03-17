package services

import (
	"context"
	"encoding/json"

	"entdemo/ent"
	"entdemo/ent/earthquake"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func ConvertStructIdentifyToNotStruct(myStruct any) struct{} {
	jsonData, err := json.Marshal(myStruct)
	if err != nil {

		return struct{}{}
	}
	jsonStr := string(jsonData)
	var unknownStruct struct{}
	err = json.Unmarshal([]byte(jsonStr), &unknownStruct)
	return unknownStruct
}

func ConvertIntToTimeStamp(milliseconds int64) time.Time {
	baseTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	// Convert milliseconds to a duration
	duration := time.Duration(milliseconds) * time.Millisecond
	// Add the duration to the base time
	timestamp := baseTime.Add(duration)
	return timestamp

}

func connect_db_v2() (*ent.Client, error) {
	client, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=root dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
	}

	return client, err
}
func CheckError(description string, err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}

func Service_get_db() []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	CheckError("Database error:", err)
	var earthquakes, errEarth = client.Earthquake.Query().Limit(100).All(ctx)
	CheckError("Background initialize error:", errEarth)
	return earthquakes
}

func Service_get_db_by_paging(pageIndex int, pageSize int) []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	CheckError("Database error:", err)
	var earthquakes, errEarth = client.Earthquake.Query().Limit(pageSize).Offset((pageIndex - 1) * pageSize).Clone().Order(func(s *sql.Selector) {
		sql.OrderByField("updated_time", sql.OrderDesc())
	}).All(ctx)
	CheckError("Background initialize error:", errEarth)
	return earthquakes
}

func Service_filter_by_status(status string) []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	CheckError("Database error:", err)

	var earthquakes, errEarth = client.Earthquake.Query().
		Where(earthquake.Status(status)).
		Order(ent.Desc("time")).All(ctx)

	CheckError("Background initialize error:", errEarth)
	return earthquakes
}
