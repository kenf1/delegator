package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kenf1/delegator/src/db"
	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/models"
	"github.com/kenf1/delegator/src/test"
)

func ReadAllTasks(w http.ResponseWriter, r *http.Request) {
	io.TasksMutex.RLock()
	defer io.TasksMutex.RUnlock()

	//copy to prevent race condition
	res := make([]models.TaskDBRow, len(test.Tasks))
	copy(res, test.Tasks)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadSingleTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	io.TasksMutex.RLock()
	defer io.TasksMutex.RUnlock()

	task, _, entryPresent := db.FindTaskByID(test.Tasks, id)
	if !entryPresent {
		http.Error(w, "entry not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
