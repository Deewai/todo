build:
	cd todo-service && $(MAKE) build
	cd todo-client && $(MAKE) build

run: 
	cd todo-service && SERVICE_PORT=${SERVICE_PORT} SERVICE_NAME=todoService $(MAKE) run
	cd todo-client && CLIENT_PORT=${CLIENT_PORT} SERVICE_NAME=todoClient SERVER_ADDRESS=todoService:${SERVICE_PORT} $(MAKE) run
	docker network rm todoNetwork
	docker network create todoNetwork
	docker network connect todoNetwork todoService
	docker network connect todoNetwork todoClient
	echo Client is running on port localhost:${CLIENT_PORT}