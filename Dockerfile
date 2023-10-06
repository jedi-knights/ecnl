# Use an official Golang runtime as a parent image
FROM golang:1.17

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Copy and rename the YAML configuration file into the container
COPY config/prod-config.yaml /app/config.yaml

# Build the Go application inside the container
RUN go build -o ecnl

# Expose the port that your Go Echo API will listen on
EXPOSE 8080

# Define the command to run your application
CMD ["./ecnl", "api"]
