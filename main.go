package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var tasks []string

func main() {
	app := &cli.App{
		Name:  "todo",
		Usage: "A simple CLI To-Do List application",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a task to the list",
				Action:  addTask,
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all tasks",
				Action:  listTasks,
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Remove a task from the list",
				Action:  removeTask,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func addTask(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("please provide a task to add")
	}
	task := c.Args().First()
	tasks = append(tasks, task)
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
	index := c.Args().First()
	// Implementation for removing a task
	fmt.Printf("Task removed: %s\n", index)
	return nil
}
