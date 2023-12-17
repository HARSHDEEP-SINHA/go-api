package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
type TaskList struct {
	sync.RWMutex
	Tasks []Task
}

var tasks = TaskList{Tasks: make([]Task, 0)}

func main() {

	http.HandleFunc("/api/tasks", getTasksHandler)
	http.HandleFunc("/api/tasks/add", addTaskHandler)
	http.HandleFunc("/api/tasks/complete", completeTaskHandler)
	fmt.Println("Server is listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks.RLock()
	defer tasks.RUnlock()

	jsonResponse, err := json.Marshal(tasks.Tasks)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	tasks.Lock()
	defer tasks.Unlock()

	newTask.ID = len(tasks.Tasks) + 1

	tasks.Tasks = append(tasks.Tasks, newTask)

	jsonResponse, err := json.Marshal(newTask)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func completeTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskToUpdate Task
	err := json.NewDecoder(r.Body).Decode(&taskToUpdate)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	tasks.Lock()
	defer tasks.Unlock()
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == taskToUpdate.ID {
			tasks.Tasks[i].Completed = true
			jsonResponse, err := json.Marshal(tasks.Tasks[i])
			if err != nil {
				http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")

			w.Write(jsonResponse)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
