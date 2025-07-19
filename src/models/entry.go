package models

type TaskEntry struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}
