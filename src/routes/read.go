package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kenf1/delegator/src/models"
)

func ReadTasks(w http.ResponseWriter, r *http.Request) {
	taskEntry := []models.TaskEntry{
		{Id: 1, Task: "Spin up service", Status: "running"},
		{Id: 2, Task: "Connect mainframe", Status: "queued"},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(taskEntry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
