package routes

import (
	"fmt"
	"net/http"

	"github.com/kenf1/delegator/src/db"
	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/models"
	"github.com/kenf1/delegator/src/test"
)

func deleteByIndex(tasks []models.TaskDBRow, index int) ([]models.TaskDBRow, error) {
	if index < 0 || index >= len(tasks) {
		return tasks, fmt.Errorf("index %d out of range", index)
	}
	return append(tasks[:index], tasks[index+1:]...), nil
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	io.TasksMutex.Lock()
	defer io.TasksMutex.Unlock()

	_, index, entryPresent := db.FindTaskByID(test.Tasks, id)
	if !entryPresent {
		http.Error(w, "entry not found", http.StatusBadRequest)
		return
	}

	var err error
	test.Tasks, err = deleteByIndex(test.Tasks, index)
	if err != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
