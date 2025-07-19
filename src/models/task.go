package models

type TaskDBRow struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}
