build:
	# protoc --proto_path=$(GOPATH)/src/github.com/Deewai/todo/proto --go_out=plugins=grpc:./internal/proto/todo  $(GOPATH)/src/github.com/Deewai/todo/proto/todo.proto
	docker run -v $(GOPATH)/src/github.com/Deewai/todo:/defs namely/protoc-all -f proto/todo.proto -l go -o ./todo-client/internal/proto/todo
	docker build -t todo-client .

run:
	docker run -e CLIENT_PORT -e SERVER_ADDRESS -p ${CLIENT_PORT}:${CLIENT_PORT} todo-client


test:
	