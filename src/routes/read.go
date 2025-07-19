package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/models"
	"github.com/kenf1/delegator/src/test"
)

func ReadAllTasks(w http.ResponseWriter, r *http.Request) {
	io.TasksMutex.RLock()
	defer io.TasksMutex.RUnlock()

	//copy to prevent race condition
	res := make([]models.TaskEntry, len(test.Tasks))
	copy(res, test.Tasks)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadSingleTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	io.TasksMutex.RLock()
	defer io.TasksMutex.RUnlock()

	var task models.TaskEntry
	found := false
	for _, t := range test.Tasks {
		if t.Id == id {
			task = t
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
