<h1>CLI To-Do List Manager</h1>



<!-- Program overview -->
<p>This program allows a user to manage a todo list from the command line. Users can add a tasks, delete a task, clear all tasks, and exit the program. Tasks are stored on the system in the file "tasks.txt" and in the program as a slice. User input is read using bufio and manipulates the task list using a switch statement. The Task type is a struct with an ID and name. The program outputs the current task list and available commands to the user.</p>



<!-- Prerequisites -->
<h2><b>Prerequisites</b></h2>
<p>To run this program, you will need Go installed on your system. You can download it from:</p>

[golang.org](https://go.dev/)



<!-- Installation -->
<h2><b>Installation</b></h2>

- Clone the repository.
- (Optional) In the terminal, traverse to the project directory.
- Run the following command: go run todolist.go.
  - If terminal's working directory is not the project directory, replace filename 'todolist.go' with file path.



<!-- How It Works -->
<h2><b>How It Works</b></h2>
<p>Upon running the program, the user will be presented with the current task list (or an empty task list, if the task list is empty) and the available options. To select an option, the user should enter the corresponding number and return. The program will then execute the selected option and loop back to the beginning.</p>

![OptionsMenu_EmptyList](https://user-images.githubusercontent.com/96446640/236118835-15f89bcd-f41d-4ced-999b-8af934860fe7.png)



<!-- Adding a Task -->
<h2><b>Adding a Task</b></h2>
<p>To add a new task to the list, select option 1 and enter a task description when prompted.</p>

- Does not add task to the list if an empty task description is entered.
- Attempts to overwrite "tasks.txt" with an updated task list.

![AddingTask](https://user-images.githubusercontent.com/96446640/236120285-332f27bb-3619-4d88-9aa2-14d92dea6826.png)



<!-- Deleting a Task -->
<h2><b>Deleting a Task</b></h2>
<p>To delete a task from the list, select option 2 and enter the task ID when prompted.</p>

- Checks if there are any tasks to delete.
- Iterates over the task list to check if the task ID exists.
  - If it exists, remove the task from the list and reorder the task ID.
- If successful, attempt to overwrite "tasks.txt" file with an updated task list.

![DeletingTask](https://user-images.githubusercontent.com/96446640/236121606-610071be-1d89-4f92-824a-d06248e8807f.png)



<!-- Clearing the Task List -->
<h2><b>Clearing the Task List</b></h2>
<p>To clear all tasks from the list, select option 3.</p>

- Attempts to overwrite "task.txt" file with an empty task list.

![ClearingTaskList](https://user-images.githubusercontent.com/96446640/236123065-76a62452-7914-4940-93f3-0d50237cc23b.png)



<!-- Exiting the Program -->
<h2><b>Exiting the Program</b></h2>
<p>To exit the program, select option 4.</p>

![ExitingProgram](https://user-images.githubusercontent.com/96446640/236123596-a3b7379e-5ade-4714-a5d6-bdf452a558e8.png)
