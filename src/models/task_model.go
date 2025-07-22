package models

type TaskDBRow struct {
	Id     string `json:"id"` //uuid string
	Task   string `json:"task"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

type PatchRequest struct {
	Id     string  `json:"id"` //required
	Task   *string `json:"task,omitempty"`
	Status *string `json:"status,omitempty"`
}
