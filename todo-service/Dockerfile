FROM golang:latest


LABEL maintainer="Innocent Abdullahi <deewai48@gmail.com>"

RUN apt-get update


# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/Deewai/todo/todo-service

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
ENV GO111MODULE=on
RUN go mod download
# RUN go mod vendor

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .
RUN go install -v ./

# Run the executable
CMD ["todo-service"]
