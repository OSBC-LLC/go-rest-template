# Orchestration Rest micro service template written in Go
- Go REST micro service layout with go-chi router, internal controllers, and pkg services
- Internal config to connect to a Postgres database as well as runtime variables pulled from env
- Unit tests on config and all routes
- Docker setup with docker-compose
- Nginx folder with basic config ( needs tweaking and updating to actually work with service ) ( not tested in heroku yet )
- Pre-commit hooks
- Jenkins with enterprise pipeline integration
- Makefile of common commands
- init.sh script to run and replace all template module names to be your project’s base module. Instructions in Readme
- Swagger doc generator
- HMAC checks on routes

## Prerequisites
- [pre-commit cli](https://pre-commit.com/)
- [docker](https://www.docker.com/products/docker-desktop/)
- [golang](https://go.dev/)
- [golang-ci](https://github.com/golangci/golangci-lint)
- [ripgrep](https://github.com/BurntSushi/ripgrep)
- [ginkgo cli](https://onsi.github.io/ginkgo/#getting-started)

## Common Docs / Links
- [Gorm ORM](https://gorm.io/docs/)
- [Go-Chi Router](https://github.com/go-chi/chi)
- [Ginkgo](https://onsi.github.io/ginkgo/)

## Initial Setup
Take a look at these files and do any necessary changes to paths, names, or config details.

- go.mod
- Dockerfile
- docker-compose.psql.yml
- docker-compose.yml
- sonar-project.properties
- Procfile
- Makefile

Run the **init.sh** script to enter in a new module name and make the replacement in all files.
```
./init.sh

# enter the name of the module:
github.com/demoOrg/repo-name

# you may need to remove other files that are not needed after the script runs
rm **/**-e
```
This project template already has two routes: **/heartbeat** and **/todo**. It's recommended to keep the heartbeat route, but feel free to edit and change the example todo route.

## Common Commands
These are common commands to start up docker images, build code, run pre-commit hooks, and more.

**Build project binary to bin/**
```make build```

**Testing commands**
```
# Use Makefile to run test
make test

# Create new test file
make newtest
# enter the name of the test. (ex heartbeat will create a heartbeat_test.go file in tests/)
heartbeat
# new test generated.

# See coverage report locally as html
make report

# Run tests
go test -v ./tests -coverprofile=./coverage.out -coverpkg ./...

# See coverage report
go tool cover -html=coverage.out

# Generate coverate and test reports for SonarQube
go test ./tests -coverprofile=coverage.out -coverpkg ./... -json > test-report.out
```

**Docker commands**
```
# Start only postgres database
make dp

# Build all docker images
make db

# Start all docker images
make du

# Run both db & du commands
make da
```

**Pre-commit commands**
```
#Install all pre-commit hooks
pre-commit install

# Force run all pre-commit hooks
pre-commit run --all-files
```
