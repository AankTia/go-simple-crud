package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AankTia/go-simple-crud/database"
	models "github.com/AankTia/go-simple-crud/models"
	"github.com/gorilla/mux"
)

// Response is a generic response struct
type Response struct {
	Status 	int 		`json:"status"`
	Message string 		`json:"message"`
	Data 	interface{} `json:"data,omitempty"`
}

// GetAllTasks return all tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	result := database.DB.Find(&tasks)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving tasks")
		return
	}
}

// GetTask returns a single task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	result := database.DB.First(&task, params["id"])
	if result.Error != nil {
		respondWithError(w, http.StatusNotFound, "Task not found")
		return
	}

	respondWithJSON(w, http.StatusOK, Response{
		Status: http.StatusOK,
		Message: "Task retrieview sucessfully",
		Data: task,
	})
}

// CreateTask creates a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	decoder := json.NewDecoder(r.Body)

	if err:= decoder.Decode(&task); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	result := database.DB.Create(&task)
	if result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating task")
		return
	}

	respondWithJSON(w, http.StatusCreated, Response{
		Status: http.StatusCreated,
		Message: "Task created successfully",
		Data: task,
	})
}

// UpdateTask updates an existing task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	// Check if task is exists
	result := database.DB.First(&task, params["id"])
	if result.Error != nil {
		respondWithError(w, http.StatusNotFound, "Tas not foun")
		return
	}

	// Dcode the updated task
	var updatedTask models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedTask); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Update field
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status

	// Save changes
	respondWithJSON(w, http.StatusOK, Response{
		Status: http.StatusOK,
		Message: "Task updated successfully",
		Data: task,
	})
}

// DeleteTask deletes a task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	// Check if task exists
	result := database.DB.First(&task, params["id"])
	if result.Error != nil {
		respondWithError(w, http.StatusNotFound, "Task not found")
	}

	// Delete the task
	database.DB.Delete(&task)

	respondWithJSON(w, http.StatusOK, Response{
		Status: http.StatusOK,
		Message: "Task deleted successfully",
	})
}

// Helper function to responsd with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Helper function to respond with an error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, Response{
		Status: code,
		Message: message,
	})
}