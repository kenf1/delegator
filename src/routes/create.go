package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/models"
	"github.com/kenf1/delegator/src/test"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var reqBody models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	newTask := models.TaskDBRow{
		Id:     123,
		Task:   reqBody.Task,
		Status: reqBody.Status,
	}

	io.TasksMutex.Lock()
	defer io.TasksMutex.Unlock()

	test.Tasks = append(test.Tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
