# Go_CLI_ToDo


<!-- Program overview -->
This program allows a user to manage a todo list from the command line. Users can add a tasks, delete a task, clear all tasks, and exit the program. Tasks are stored on the system in the file "tasks.txt" and in the program as a slice. User input is read using bufio and manipulates the task list using a switch statement. The Task type is a struct with an ID and name. The program outputs the current task list and available commands to the user.

<br>

<!-- Prerequisites -->
<h2>Prerequisites</h2>
To run this program, you will need Go installed on your system. You can download it from:

[golang.org](https://go.dev/)

<br>


<!-- Installation -->
<h2>Installation</h2>
<ol>
  <li>
    Clone the repository.
  </li>
  <li>
    (Optional) In the terminal, traverse to the project directory.
  </li>
  <li>
    Run the following command: go run todolist.go
  </li>
  <ul>
    <li>
      If terminal's working directory is not the project directory, replace filename 'todolist.go' with file path.
    </li>
  </ul>
</ol>

<br>


<!-- Usage -->
<h2>Usage</h2>
Upon running the program, the user will be presented with the current task list (or an empty task list, if the task list is empty) and the available options:

![OptionsMenu_EmptyList](https://user-images.githubusercontent.com/96446640/236118835-15f89bcd-f41d-4ced-999b-8af934860fe7.png)

To select an option, the user should enter the corresponding number and return. The program will then execute the selected option and loop back to the beginning.

<br>


<!-- Adding a Task -->
<h2>Adding a Task</h2>
To add a new task to the list, select option 1 and enter a task description when prompted:

![AddingTask](https://user-images.githubusercontent.com/96446640/236120285-332f27bb-3619-4d88-9aa2-14d92dea6826.png)

<ul>
  <li>
    Does not add task to the list if an empty task description is entered.
  </li>
</ul>

<br>


<!-- Deleting a Task -->
<h2>Deleting a Task</h2>
To delete a task from the list, select option 2 and enter the task ID when prompted:

![DeletingTask](https://user-images.githubusercontent.com/96446640/236121606-610071be-1d89-4f92-824a-d06248e8807f.png)

<ul>
  <li>
    Checks if there are any tasks to delete.
  </li>
  <li>
    Iterates over the task list to check if the task ID exists.
  </li>
  <ul>
    <li>
      If it exists, remove the task from the list and reorder the task ID
    </li>
  </ul>
</ul>
