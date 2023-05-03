/*
This program allows a user to manage a todo list from the command line. Users can add a task, delete a task, clear all tasks,
and exit the program. Tasks are stored on the sytem, in the file "tasks.txt", and in the program as a slice. User input is read
using bufio and manipulates the taks list using a switch statement. The Task type is a struct with an ID and name. The program
outputs the current task list and available commands to the user.
To run this program from the command line, use the instruction:		go run todolist.go
Replace the file name with the file path, if you are not in the file's directory.
*/

package main

import (
	"bufio"   // Provides buffered I/O operations
	"fmt"     // Provides formatted I/O
	"os"      // Provides platform-independent OS functions
	"strconv" // Provides functions for string conversion
	"strings" // Provides functions for manipulating strings
)

// Represents a task in the todo list
type Task struct {
	id   int    // Unique identifier for the task
	name string // Name or description of the task
}

// The main function that handles user input and outputs task list options
func main() {
	tasks := loadTasks()                // Load tasks from file into a map
	reader := bufio.NewReader(os.Stdin) // Initialize reader for user input

	// Infinite loop
	for {
		// Checks whether there are tasks, then either print a message, or print all tasks
		if len(tasks) == 0 {
			fmt.Println("Todo List: empty")
		} else {
			fmt.Println("Todo List:")
			for _, task := range tasks {
				fmt.Printf("%d. %s\n", task.id, task.name)
			}
		}

		// Print available options
		fmt.Println("\nOptions:")
		fmt.Println("1. Add a task")
		fmt.Println("2. Delete a task")
		fmt.Println("3. Clear all tasks")
		fmt.Println("4. Exit")
		fmt.Print("Enter an option: ")

		// Read user input and remove whitespace
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		fmt.Println()

		// Switch statement to handle different user options
		switch option {
		case "1": // Add a task
			// Prompt the user to enter a tasks, and store the input without leading/trailing whitespace
			fmt.Print("Enter a task: ")
			var name string
			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)

			// Check input length, and break if length is zero
			if len(name) == 0 {
				fmt.Println("No task was entered.")
				break
			}

			// Create a new Task object with the next available ID, and append it to tasks
			task := Task{id: len(tasks) + 1, name: name}
			tasks = append(tasks, task)

			saveTasks(tasks) // Calls saveTasks function to create a file or overwrite the file with the current tasks

			fmt.Printf("Added task '%s' with id %d\n", task.name, task.id)
		case "2": // Delete a task
			// Check whether there are no tasks. If so, print an error message and exit the switch
			if len(tasks) == 0 {
				fmt.Println("The todo list is empty.")
				break
			}

			// Prompt the user to enter an ID, remove the leading/trailing whitespace, and store the input as an integer
			fmt.Print("Enter task id: ")
			var idStr string
			idStr, _ = reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			found := false
			// Iterate over the tasks
			for index, task := range tasks {
				// Check whether the ID of the current task is the ID to be deleted
				if task.id == id {
					found = true
					tasks = append(tasks[:index], tasks[index+1:]...) // Create a new slice, excluding the removed element
					fmt.Printf("Deleted task with id %d\n", id)
					break
				}
			}

			// Check whether the ID was found. If so, update the task's ID. Otherwise, print an error message
			if found {
				for i := range tasks {
					tasks[i].id = i + 1
				}
				saveTasks(tasks) // Calls saveTasks function to overwrite the file with the current tasks
			} else {
				fmt.Println("Task does not exist.")
			}
		case "3": // Clear the task list
			tasks = []Task{} // overwrites task with an empty slice of type Task
			fmt.Println("Cleared all tasks")
			saveTasks(tasks) // Calls saveTasks function to overwrite the file with the no tasks
		case "4": // Exit
			fmt.Println("Goodbye!")
			return
		default: // Invalid option
			fmt.Println("Invalid option.")
		}

		// print a separator between the current and next iteration
		fmt.Println("--------------------------------------------------")
	}
}

/*
Reads tasks from a file named "tasks.txt" and creates a slice of Task objects from the file contents
Returns []Task: a slice of Task objects read from the file
*/
func loadTasks() []Task {
	file, err := os.OpenFile("tasks.txt", os.O_RDONLY|os.O_CREATE, 0644)

	// Panic if there was an error opening the file
	if err != nil {
		panic(err)
	}

	defer file.Close()                // Defers closing the file until the function returns
	scanner := bufio.NewScanner(file) // Creates a scanner to read from the file
	var tasks []Task                  // Creates an empty slice of tasks

	// Reads the file line by line
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "|") // Splits each line into fields using the pipe '|' separator
		id, err := strconv.Atoi(fields[0])           // Converts the first field, the task ID, from a string to an integer

		// Panic if there was an error converting ID to an integer
		if err != nil {
			panic(err)
		}

		task := Task{id: id, name: fields[1]} // Creates a new task struct with the ID and name read from the file
		tasks = append(tasks, task)           // Adds the task to the slice
	}

	return tasks // Returns the slice of tasks
}

/*
Saves a slice of tasks to a file name "tasks.txt"
Parameter []Task: a slice of Task objects to write to the file
*/
func saveTasks(tasks []Task) {
	file, err := os.Create("tasks.txt") // Create a new file named "tasks.txt" and open it for writing

	// Panic if there was an error opening the file
	if err != nil {
		panic(err)
	}

	defer file.Close()              // Defer closing the file until the function returns
	writer := bufio.NewWriter(file) // Create a new buffered writer that will write to the file

	// Iterate over each task in the slice
	for _, task := range tasks {
		// Write the task ID and name to the file in the format "ID|name\n"
		_, err = fmt.Fprintf(writer, "%d.%s\n", task.id, task.name)

		// Panic if there was an error writing to the file
		if err != nil {
			panic(err)
		}
	}

	writer.Flush() // Flush any remaining data in the writer to the file
}
