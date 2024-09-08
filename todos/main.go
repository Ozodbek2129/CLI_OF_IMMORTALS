package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

var tasks []string

const tasksFile = "tasks.json"

func main() {
	loadTasks()

	app := &cli.App{
		Name:  "todo",
		Usage: "A simple CLI todo list manager",
		Commands: []*cli.Command{
			{
				Name:   "add",
				Usage:  "Add a new task",
				Action: addTask,
			},
			{
				Name:   "list",
				Usage:  "List all tasks",
				Action: listTasks,
			},
			{
				Name:   "remove",
				Usage:  "Remove a task",
				Action: removeTask,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func loadTasks() {
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Tasks file does not exist yet")
		} else {
			fmt.Println("Error reading tasks file:", err)
		}
		return
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks:", err)
	} else {
		fmt.Printf("Loaded %d tasks\n", len(tasks))
	}
}

func saveTasks() {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return
	}
	err = os.WriteFile(tasksFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks file:", err)
	} else {
		fmt.Println("Tasks saved successfully")
	}
}

func addTask(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("please provide a task description")
	}
	task := c.Args().First()
	tasks = append(tasks, task)
	saveTasks()
	fmt.Printf("Task added: %s\n", task)
	return nil
}

func listTasks(c *cli.Context) error {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list")
		return nil
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
	return nil
}

func removeTask(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("please provide a task number to remove")
	}
	index, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf("invalid task number: %v", err)
	}
	index-- // Adjust for 0-based indexing
	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("task number out of range")
	}
	removedTask := tasks[index]
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTasks()
	fmt.Printf("Task removed: %s\n", removedTask)
	return nil
}
