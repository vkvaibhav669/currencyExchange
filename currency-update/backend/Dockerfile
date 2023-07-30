# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container's working directory
COPY main.go .

# Build the Go application
RUN go build -o app .

# Expose the port that the Go application listens on
EXPOSE 8080

# Command to run the Go application
CMD ["./app"]
