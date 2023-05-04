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
