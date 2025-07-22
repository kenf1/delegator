package db

import "github.com/kenf1/delegator/src/models"

func FindTaskByID(tasks []models.TaskDBRow, id string) (*models.TaskDBRow, int, bool) {
	for i, t := range tasks {
		if t.Id == id {
			return &tasks[i], i, true
		}
	}
	return nil, -1, false
}
