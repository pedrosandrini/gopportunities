.PHONY: default run build test docs clean
#VARIABLES
APP_NAME=gopportunities

#TASKS
default: run-with-docs

run: 
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go4
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -f ./docs