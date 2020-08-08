build:
	# protoc --proto_path=$(GOPATH)/src/github.com/Deewai/todo/proto --go_out=plugins=grpc:./internal/proto/todo  $(GOPATH)/src/github.com/Deewai/todo/proto/todo.proto
	docker run -v $(GOPATH)/src/github.com/Deewai/todo:/defs namely/protoc-all -f proto/todo.proto -l go -o ./todo-service/internal/proto/todo
	docker build -t todo-service .
run:
	docker run -e SERVICE_PORT -d -p ${SERVICE_PORT}:${SERVICE_PORT} todo-service