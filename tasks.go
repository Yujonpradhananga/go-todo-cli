package main

import (
    "encoding/json"
    "fmt"
    "os"
    "strconv"
)

type Task struct {
    Description string `json:"description"`
    Done        bool   `json:"done"`
}

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
    file, err := os.ReadFile(fileName)
    if err != nil {
        return []Task{}, nil
    }

    var tasks []Task
    err = json.Unmarshal(file, &tasks)
    return tasks, err
}

func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(fileName, data, 0644)
}

func AddTask(description string) {
    tasks, _ := loadTasks()
    tasks = append(tasks, Task{Description: description})
    saveTasks(tasks)
    fmt.Println("Added:", description)
}

func ListTasks() {
    tasks, _ := loadTasks()
    for i, task := range tasks {
        status := "[ ]"
        if task.Done {
            status = "[x]"
        }
        fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
    }
}

func MarkDone(indexStr string) {
    index, err := strconv.Atoi(indexStr)
    if err != nil {
        fmt.Println("Invalid task number.")
        return
    }

    tasks, _ := loadTasks()
    if index <= 0 || index > len(tasks) {
        fmt.Println("Task number out of range.")
        return
    }

    tasks[index-1].Done = true
    saveTasks(tasks)
    fmt.Println("Marked as done:", tasks[index-1].Description)
}
