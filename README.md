The provided Go project demonstrates a simple HTTP server that serves as a basic task tracker API. Below is an overview of the project structure and functionality:

Project Structure:
main.go:

The main entry point of the application.
Defines the main function that sets up HTTP routes and starts the server.
Task Struct:

Represents a simple task structure with fields for ID, description, and completion status.
TaskList Struct:

Represents a list of tasks with a read-write mutex for concurrent access control.
Functionality:
HTTP Server Setup:

The main function sets up the HTTP server to listen on port 8080.
API Endpoints:

/api/tasks (GET):
Returns the list of tasks in JSON format.
/api/tasks/add (POST):
Adds a new task to the list. Expects a JSON payload with a task description.
/api/tasks/complete (POST):
Marks a task as completed. Expects a JSON payload with the task ID.
Task Handling:

The tasks are stored in-memory within the tasks variable, which is an instance of the TaskList struct.
The sync.RWMutex is used to control access to the task list, allowing multiple readers or a single writer at a time.
Request Handling:

Request handlers (getTasksHandler, addTaskHandler, and completeTaskHandler) process incoming HTTP requests.
These handlers perform tasks such as retrieving the task list, adding a new task, and marking a task as completed.
They use the encoding/json package to handle JSON encoding and decoding.
Server Start and Stop:

The http.ListenAndServe function starts the server.
The server can be stopped manually by pressing Ctrl+C in the terminal.
How to Test:
Run the Program:

Execute go run main.go in the terminal to start the server.
Test API Endpoints:

Use tools like curl, a web browser, or Postman to interact with the API.
Examples: curl http://localhost:8080/api/tasks, curl -X POST -H "Content-Type: application/json" -d '{"description": "Buy groceries"}' http://localhost:8080/api/tasks/add
Verify Responses:

Check the terminal for server logs and inspect the responses from API requests.
Stop the Server:

Manually stop the server by pressing Ctrl+C in the terminal.
Notes:
This is a basic example for educational purposes and may not be suitable for production.
In a real-world scenario, additional features, error handling, and security measures would be necessary.
The tasks are stored in-memory and will be lost when the server is stopped.
 To run the above Go code, follow these steps:

Install Go:
Make sure you have Go installed on your system. You can download and install Go from the official website: https://golang.org/dl/

Create a file for the code:
Copy the provided Go code into a file, for example, main.go.

Open a terminal or command prompt:
Open a terminal or command prompt in the directory where you saved the main.go file.

Run the code:
Execute the following command to run the Go program:

bash
Copy code
go run main.go
This command will compile and run the Go code. If everything is set up correctly, you should see the message "Server is listening on :8080" in the terminal.

Test the API:
Open a web browser or use a tool like curl to test the API. For example, you can use curl to get the list of tasks:

bash
Copy code
curl http://localhost:8080/api/tasks
You can also use other HTTP clients or tools, such as Postman, to interact with the API.

Remember that this is a simple in-memory example, and the data won't persist between program executions. In a real-world scenario, you would likely use a database for data storage.

To stop the running server, you can typically press Ctrl+C in the terminal where the server is running.

Feel free to reach out if you encounter any issues or have further questions!





