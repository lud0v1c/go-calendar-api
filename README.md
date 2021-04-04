# Go Calendar API

## About
This is a back-end only implementation of a simple calendar application to schedule slots between candidates and interviewers. Done as a challenge for a technical interview, it uses entirely Go with the Gin framework and standard RESTful practices.

## Structure
```bash
├── api
│   ├── controllers.go # "Middleware"-like handlers
│   ├── database.go # All BD operations
│   ├── models.go # Models of our entities
│   └── routers.go # Routing, sets the endpoints
├── go.mod # Module file
├── gorm.db # sqlite DB used by gorm
├── go.sum # Module file
├── main.go # Just a small stub to launch the server
```

This was developed has a [Go Module](https://golang.org/ref/mod), so it can be easily imported into other projects.

## Installation
The project was developed and tested with **go1.16.2 linux/amd64**, but a Dockerfile is provided so it can be deployed quickly without needing to install Go and setup a Workspace.
### Building & Running with Go
1) Install the latest Go for [your desired platform](https://golang.org/doc/install).
2) Setup your workspace ($GOPATH) accordingly, so you can import everything needed to run the project:
```bash
export GOPATH="/your/target/workspace/dir"
```
3) Confirm Go is working with _go version_, and run the project from this directory (no need to build it):
```bash
go run main.go
```
### Using Docker
If you have Docker installed, you can use the provided Dockerfile to build and run a containerized version of this project, without needing to install Go or anything else.

1) In the project directory, build the image:
```bash
sudo docker build -t go-calendar-api .
```
2) Run it (note that the port is hardcoded in _main.go_ - change it there too if you don't want to use port 8080):
```bash
sudo docker run -p 8080:8080 -it go-calendar-api
```
## Usage
The API endpoints can be viewed in [routers.go](api/routers.go). It follows a RESTful fashion, so the 3 main components (users, slots and scheduling) can be reached on **api/user, api/slots/ and api/scheduler**.

The operation itself might require a parameter being passed in the URI, a JSON payload in the request body, or both.

Once it's running, it can be interacted directly with any network transfer tool like Wget or curl:
```bash
curl --request POST \
        --data '{"name":"Carl", "week":33, "day": "monday", "hour_start":9, "hour_end":10}' \
  http://localhost:8080/api/slots/
curl --request GET http://localhost:8080/api/slots/Carl
```
Alternatively, any modern browser with Console/Developer tools can render JSON replies without any problem. For even fancier control, [Postman](https://www.postman.com/) is a really powerful option.
## Testing
A simple bash script, [test_example.sh](test_example.sh) is provided, simulating the scenario detailed in the [assignment](docs/assignment.md). Just _chmod +x_ it and pass as first argument the API's port number.