package services

import (
	"context"

	"entdemo/ent"
	"entdemo/model"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "root"
// 	password = "secret"
// 	dbname   = "postgres"
// 	sslmode  = "disable"
// )

const (
	host        = "localhost"
	port        = "5433"
	user        = "postgres"
	password    = "test12"
	dbname      = "postgres"
	sslmode     = "disable"
	//search_path = "simple_bank"
)

func ConvertIntToTimeStamp(milliseconds int64) time.Time {
	baseTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	// Convert milliseconds to a duration
	duration := time.Duration(milliseconds) * time.Millisecond
	// Add the duration to the base time
	timestamp := baseTime.Add(duration)
	return timestamp

}

func connect_db_v2() (*ent.Client, error) {
	client, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=root dbname=postgres password=secret sslmode=disable sslmode=disable") // search_path=simple_bank
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return client, err
}
func checkError(description string, err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}

func Service_get_db() []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	checkError("Loi db", err)
	var earthquakes, errEarth = client.Earthquake.Query().All(ctx)
	checkError("loi khoi tao background", errEarth)
	return earthquakes
}

func Service_get_db_by_paging(pageIndex int, pageSize int) []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	checkError("Loi db", err)
	//limit : lay ra bao nhieu thang , offset bo qua bao nhieu thang
	var earthquakes, errEarth = client.Earthquake.Query().Limit(pageSize).Offset((pageIndex - 1) * pageSize).Clone().All(ctx)
	checkError("loi khoi tao background", errEarth)
	return earthquakes
}

func Service_get_clause_db_by_paging(filter model.EarthquakeFilterModel) []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	checkError("Loi db", err)

	var updateTimeTo = ConvertIntToTimeStamp(int64(filter.UpdateTimeTo))
	var updateTimeFrom = ConvertIntToTimeStamp(int64(filter.UpdateTimeFrom))
	//limit : lay ra bao nhieu thang , offset bo qua bao nhieu thang
	var earthquakes, errEarth = client.Earthquake.Query().Limit(filter.PageIndex).Offset((filter.PageIndex - 1) * filter.PageSize).Where(func(s *sql.Selector) {
		s.Where(
			sql.And(
				sql.EQ("feature_id", filter.Id),
				sql.LTE("update_time", updateTimeFrom),
				sql.GTE("update_time", updateTimeTo),
			),
		)
	}).Clone().All(ctx)
	checkError("loi khoi tao background", errEarth)
	return earthquakes
}
