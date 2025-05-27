package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: todo [add|list|done] [task]")
        return
    }

    command := os.Args[1]
    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Please provide a task description.")
            return
        }
        task := os.Args[2]
        AddTask(task)

    case "list":
        ListTasks()

    case "done":
        if len(os.Args) < 3 {
            fmt.Println("Please provide the task number to mark as done.")
            return
        }
        MarkDone(os.Args[2])

    default:
        fmt.Println("Unknown command:", command)
    }
}
