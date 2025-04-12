package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
type Task struct {
	ID     int
	Title  string
	Done   bool
}

var tasks []Task
var idCounter int = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n📌 Simple To-Do List")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTask(scanner)
		case "2":
			viewTasks()
		case "3":
			markTaskDone(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Exiting... 👋")
			return
		default:
			fmt.Println("❌ Invalid choice, please try again.")
		}
	}
}

// Function to add a new task
func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := scanner.Text()

	task := Task{ID: idCounter, Title: title, Done: false}
	tasks = append(tasks, task)
	idCounter++
	fmt.Println("✅ Task added successfully!")
}

// Function to view all tasks
func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks found!")
		return
	}
	fmt.Println("\n📋 Your Tasks:")
	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
	}
}

// Function to mark a task as done
func markTaskDone(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to mark as done: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("❌ Invalid input. Enter a number.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Println("✅ Task marked as done!")
			return
		}
	}
	fmt.Println("❌ Task not found.")
}

// Function to delete a task
func deleteTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("❌ Invalid input. Enter a number.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("🗑️ Task deleted successfully!")
			return
		}
	}
	fmt.Println("❌ Task not found.")
}