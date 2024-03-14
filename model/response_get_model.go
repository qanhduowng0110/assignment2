package model

type ResponseGetModel struct {
	Page      any `json:"page"`
	Total     int `json:"total"`
	TotalPage int `json:"totalpage"`
}
