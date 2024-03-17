package model

type EarthquakeFilterModel struct {
	PageIndex      int `json:"pageIndex"`
	PageSize       int `json:"pageSize"`
	UpdateTimeTo   int `json:"occurTime"`
	UpdateTimeFrom int `json:"updateTime"`
	Id             string
}
