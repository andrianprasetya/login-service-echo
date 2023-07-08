FROM golang:1.19.0-alpine

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go module manifests
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app

# Set the command to run the executable when the container starts
CMD ["./app"]