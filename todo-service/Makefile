build:
	docker build -t todo-service .

run:
	docker run -e SERVICE_PORT -d --name ${SERVICE_NAME} todo-service

generate:
	protoc --proto_path=$(GOPATH)/src/github.com/Deewai/todo/proto --go_out=plugins=grpc:./internal/proto/todo  $(GOPATH)/src/github.com/Deewai/todo/proto/todo.proto