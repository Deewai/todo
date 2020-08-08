build:
	cd todo-service && $(MAKE) build
	cd todo-client && $(MAKE) build

run: 
	cd todo-service && CLIENT_PORT=${CLIENT_PORT} SERVER_ADDRESS=localhost:${CLIENT_PORT} $(MAKE) run
	cd todo-client && SERVICE_PORT=${SERVICE_PORT} $(MAKE) run