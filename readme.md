# APPLICATION
The application is a sample todo app that illustrates the use of gRPC in Go. It consists of a todo-service and a todo-client. The todo-client is a basic API gateway to the todo-service. communication between the todo-client and the todo-service is made with gRPC.

## Todo Client
As stated above, the todo-client is a basic api gateway for the todo-service. It implements the following endpoints
* GET /todos - This returns a list of available todos
* POST /todos - This adds a new todo and takes a payload of the form {"name" : "a name", "description" : "a description", "deadline" : "a deadline"}
* DELETE /todos/:id - This deletes todo with id = id

## Todo Service
The todo-service implements an in-memory datastore for adding, getting and deleting todos. It interacts with a client using gRPC and implements an interface of
```
service TodoService {
  rpc AddTodo(Todo) returns (AddResponse) {}
  rpc GetTodos(GetRequest) returns (GetResponse) {}
  rpc DeleteTodo(DeleteRequest) returns (DeleteResponse) {}
}
```

# Setting up
Application has been setup to use docker automatically.

Only required thing is to have docker installed.

The docker images can be built by running the following from the prject root
```
make build
```
The command above builds the 2 services todo-client and todo-service

Running the built docker images is as simple as running the following from the project root.
```
make run
```
*NOTE*:
* SERVER_PORT and CLIENT_PORT environment variables must be set before running the above command
* The client is exposed to the local system on port CLIENT_PORT, so requests can be made like
```
curl localhost:CLIENT_PORT/todos
```

Cleaning up the applications, i.e stopping and removing containers and removing the shared network can be done by running the following
```
make clean
```

NOTE: Application receives payload of application/json format for POST requests

Author: Abdullahi Innocent (deewai48@gmail.com, https://github.com/Deewai)