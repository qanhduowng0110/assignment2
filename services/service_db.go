package services

import (
	"context"
	"encoding/json"

	"entdemo/ent"
	"entdemo/model"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

// xu ly convert interface struct to struct khong xd
func ConvertStructIdentifyToNotStruct(myStruct any) struct{} {
	jsonData, err := json.Marshal(myStruct)
	if err != nil {

		return struct{}{}
	}

	// Chuỗi JSON đại diện cho struct
	jsonStr := string(jsonData)

	// Chuyển đổi chuỗi JSON thành struct không xác định
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
func checkError(description string, err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}

func Service_get_db() []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	checkError("Database error:", err)
	var earthquakes, errEarth = client.Earthquake.Query().All(ctx)
	checkError("Background initialize error:", errEarth)
	return earthquakes
}

func Service_get_db_by_paging(pageIndex int, pageSize int) []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	checkError("Database error:", err)
	var earthquakes, errEarth = client.Earthquake.Query().Limit(pageSize).Offset((pageIndex - 1) * pageSize).Clone().Order(func(s *sql.Selector) {
		sql.OrderByField("updated_time", sql.OrderDesc())
	}).All(ctx)
	checkError("Background initialize error:", errEarth)
	return earthquakes
}

func Service_get_clause_db_by_paging(filter model.EarthquakeFilterModel) []*ent.Earthquake {
	client, err := connect_db_v2()
	ctx := context.Background()
	checkError("Database error:", err)

	var updateTimeTo = ConvertIntToTimeStamp(int64(filter.UpdateTimeTo))
	var updateTimeFrom = ConvertIntToTimeStamp(int64(filter.UpdateTimeFrom))
	var earthquakes, errEarth = client.Earthquake.Query().Limit(filter.PageSize).Offset((filter.PageIndex - 1) * filter.PageSize).Where(func(s *sql.Selector) {
		s.Where(
			sql.And(
				sql.LTE("updated_time", updateTimeTo),
				sql.GTE("updated_time", updateTimeFrom),
			),
		)
	}).Clone().All(ctx)

	checkError("Background initialize error:", errEarth)
	return earthquakes
}

// hàm xử lý lưu các filter của api
func InsertLogApiRequest(filter model.ApiReqFilterModel) bool {
	client, err := connect_db_v2()
	ctx := context.Background()
	checkError("Database error:", err)

	var _, errorInsert = client.Apireq.Create().
		SetReqParam(filter.ReqParam).
		SetReqBody(ConvertStructIdentifyToNotStruct(filter.ReqBody)).
		SetReqMetadata(filter.ReqMetadata).
		SetReqHeaders(filter.ReqHeaders).
		SetReqTime(time.Now()).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).Save(ctx)
	checkError(" InsertLogApiRequest Background initialize error:", errorInsert)
	return true
}
