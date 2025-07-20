package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kenf1/delegator/src/db"
	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/models"
	"github.com/kenf1/delegator/src/test"
)

func PutTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.TaskDBRow
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	io.TasksMutex.Lock()
	defer io.TasksMutex.Unlock()

	_, index, entryPresent := db.FindTaskByID(test.Tasks, updatedTask.Id)
	if !entryPresent {
		http.Error(w, "entry not found", http.StatusBadRequest)
		return
	}

	test.Tasks[index] = updatedTask
	w.WriteHeader(http.StatusNoContent)
}

func PatchTask(w http.ResponseWriter, r *http.Request) {
	var patch models.PatchRequest
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	id := strings.TrimSpace(patch.Id)
	if id == "" {
		http.Error(w, "'id' field must be present in JSON body", http.StatusBadRequest)
		return
	}

	io.TasksMutex.Lock()
	defer io.TasksMutex.Unlock()

	taskPointer, _, entryPresent := db.FindTaskByID(test.Tasks, id)
	if !entryPresent {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	task := taskPointer

	//update fields from body
	if patch.Task != nil {
		task.Task = *patch.Task
	}
	if patch.Status != nil {
		task.Status = *patch.Status
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
