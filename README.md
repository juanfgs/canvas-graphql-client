# Canvas GraphQL instructions

## Installation

This tutorial assumes you have git , go environment and docker with docker-compose
 properly set up.

First clone the repository as follows:

```
$ git clone https://github.com/juanfgs/canvas-graphql-server.git
```


## Fetch the dependencies
cd to the directory where you downloaded the project then run

```
$ go get -v 
```
to fetch all the dependencies

## Configuration

The ```CANVAS_GRAPHQL_URL``` is defined in the env.example file, please make a copy of this file
as follows ``` cp env.example .env```

Afterwards you have to load the file with ```source .env``` to properly load the
API URL 

## Running the project

To run the project use

```
$ go run main.go
```

## Running tests
You can run all the tests in the project with the following command:
```
 go test ./... 
```

