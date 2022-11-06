<h1 align="center">
  ZCELERO
</h1>

# Setting up the project
## Pre-requisites
To run this application you need to have the following tools installed on your machine:

* [Git](https://git-scm.com)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Running Posterr
Right now the project runs with docker-compose, so to execute the project you must run the command `docker-compose up` in your terminal inside the project folder. It will start the API in port 8080.

The application contains a file called `.env`, there are the environment variables responsible for defining the operating mode of the Gin framework and also for defining the log level that will be displayed in the terminal. By default both are in `debug` mode.

# Testing
To run unit tests and integration test, you can run the command `docker exec <CONTAINER_NAME OS CONTAINER_ID> go test ./...`. Eg: `docker exec zcelero_app_1 go test ./...`

# Critique
## Scaling
Since this application is a monolith, its scalability will come at a high cost. This project has business rules separate from its controllers. Because of this, we can separate the services into microservices to support an incremental volume of requests. The API would be available in a BFF format for the client that wants to connect

# Comments
## Error messages
If an error occurs in the execution of the application, it will return the original error message to the client, except for validation errors. This was done to bring a Lean solution to the application, but the application creator understands that in a production ready application, the error message must be clear and designed so that the audience that will use the API can understand in a way clear what happened and how to solve the problem.

## Logging
The application creator understands that logs can be expensive for applications that process a high volume of data. Because of that, only essential data should be logged, to ensure the proper functioning of the application.

In this application, the logs are in debug mode because the exercise asked to output for execution logs. However, in an application with a high volume of events, leaving the log level in debug mode can create high costs for the company, and must be kept in debug mode. Warning or at most in info mode, but with a reduced amount of logs

## Helpers
The application's creator chose to keep some native language functions separate in a Helper to avoind the creation of new interfaces only to mock the functions in the unit tests. This decision was made to reduce the application boilerplate, to maintain a Lean solution and because the unit tests of these functions already guarantee their operation.