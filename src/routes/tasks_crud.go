package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/kenf1/delegator/src/auth"
	"github.com/kenf1/delegator/src/configs"
	"github.com/kenf1/delegator/src/db"
	"github.com/kenf1/delegator/src/models"
)

// CreateTask
//
//	@Summary		Create new task
//	@Description	Create a new task from JSON body. Returns the created task row.
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			TaskRequest	body	models.TaskRequest	true	"Task creation request body"
//	@Success		201	{object}	models.TaskDBRow	"Successfully created task"
//	@Failure		400	{string}	string	"Invalid request body"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var reqBody models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	newTask := models.TaskDBRow{
		Id:     uuid.NewString(),
		Task:   reqBody.Task,
		Status: reqBody.Status,
	}

	configs.TasksMutex.Lock()
	defer configs.TasksMutex.Unlock()

	db.Tasks = append(db.Tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ReadAllTasks
//
//	@Summary		Retrieve all existing tasks
//	@Description	Returns a list of all tasks currently stored.
//	@Tags			Tasks
//	@Produce		json
//	@Success		200	{array}		models.TaskDBRow	"List of tasks"
//	@Failure		500	{string}	string				"Internal server error"
//	@Router			/tasks [get]
func ReadAllTasks(w http.ResponseWriter, r *http.Request) {
	configs.TasksMutex.RLock()
	defer configs.TasksMutex.RUnlock()

	//copy to prevent race condition
	res := make([]models.TaskDBRow, len(db.Tasks))
	copy(res, db.Tasks)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ReadSingleTask
//
//	@Summary		Get a single task by ID
//	@Description	Fetches a task entry by its unique ID from the database.
//	@Tags			Tasks
//	@Produce		json
//	@Param			id	path		string	true	"Task ID"
//	@Success		200	{object}	models.TaskDBRow	"Task found and returned successfully"
//	@Failure		400	{string}	string	"Entry not found or invalid ID"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/tasks/{id} [get]
func ReadSingleTask(w http.ResponseWriter, r *http.Request) {
	id, err := auth.SanitizeQueryParam(r.PathValue("id"))
	if err != nil {
		http.Error(w, "entry not found", http.StatusBadRequest)
		return
	}

	configs.TasksMutex.RLock()
	defer configs.TasksMutex.RUnlock()

	task, _, entryPresent := db.FindTaskByID(db.Tasks, id)
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

// PutTask
//
//	@Summary		Update an existing task
//	@Description	Updates a task by ID using JSON body data.
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			TaskDBRow	body		models.TaskDBRow	true	"Updated task data with existing ID"
//	@Success		204	"Task updated successfully; no content returned"
//	@Failure		400	{string}	string	"Invalid JSON body or task not found"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/tasks [put]
func PutTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.TaskDBRow
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	configs.TasksMutex.Lock()
	defer configs.TasksMutex.Unlock()

	_, index, entryPresent := db.FindTaskByID(db.Tasks, updatedTask.Id)
	if !entryPresent {
		http.Error(w, "entry not found", http.StatusBadRequest)
		return
	}

	db.Tasks[index] = updatedTask
	w.WriteHeader(http.StatusNoContent)
}

// PatchTask
//
//	@Summary		Partially update an existing task
//	@Description	Update one or more fields of a task by providing JSON with the task ID and fields to change.
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			patch	body		models.PatchRequest	true	"Partial task update request including task ID"
//	@Success		200		{object}	models.TaskDBRow	"Updated task returned in response"
//	@Failure		400		{string}	string	"Invalid JSON body or missing 'id' field"
//	@Failure		404		{string}	string	"Task not found"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/tasks [patch]
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

	configs.TasksMutex.Lock()
	defer configs.TasksMutex.Unlock()

	taskPointer, _, entryPresent := db.FindTaskByID(db.Tasks, id)
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

func deleteByIndex(tasks []models.TaskDBRow, index int) ([]models.TaskDBRow, error) {
	if index < 0 || index >= len(tasks) {
		return tasks, fmt.Errorf("index %d out of range", index)
	}
	return append(tasks[:index], tasks[index+1:]...), nil
}

// DeleteTask
//
//	@Summary		Delete a task by ID
//	@Description	Deletes the task identified by the given ID.
//	@Tags			Tasks
//	@Produce		json
//	@Param			id		path		string	true	"Task ID to delete"
//	@Success		204		"Task deleted successfully; no content returned"
//	@Failure		400		{string}	string	"Entry not found or invalid ID"
//	@Failure		500		{string}	string	"Failed to delete task due to server error"
//	@Router			/tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := auth.SanitizeQueryParam(r.PathValue("id"))
	if err != nil {
		http.Error(w, "entry not found", http.StatusBadRequest)
		return
	}

	configs.TasksMutex.Lock()
	defer configs.TasksMutex.Unlock()

	_, index, entryPresent := db.FindTaskByID(db.Tasks, id)
	if !entryPresent {
		http.Error(w, "entry not found", http.StatusBadRequest)
		return
	}

	var err1 error
	db.Tasks, err1 = deleteByIndex(db.Tasks, index)
	if err1 != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func TasksRoutes() *http.ServeMux {
	tasksMux := http.NewServeMux()

	tasksMux.HandleFunc("GET /get", ReadAllTasks)
	tasksMux.HandleFunc("GET /get/{id}", ReadSingleTask)
	tasksMux.HandleFunc("POST /create", CreateTask)
	tasksMux.HandleFunc("DELETE /delete/{id}", DeleteTask)
	tasksMux.HandleFunc("PUT /put", PutTask)
	tasksMux.HandleFunc("PATCH /patch", PatchTask)

	return tasksMux
}
