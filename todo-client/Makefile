build:
	docker build -t todo-client .

run:
	docker run -e CLIENT_PORT -e SERVER_ADDRESS -d --name ${SERVICE_NAME} -p ${CLIENT_PORT}:${CLIENT_PORT} todo-client

generate:
	protoc --proto_path=$(GOPATH)/src/github.com/Deewai/todo/proto --go_out=plugins=grpc:./internal/proto/todo  $(GOPATH)/src/github.com/Deewai/todo/proto/todo.proto