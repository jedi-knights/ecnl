# Build Stage
FROM golang:1.21.2 AS BuildStage
LABEL stage=build

# These environment variables are required for the Go build.
# They tell the Go compiler to build a static binary with no CGO dependencies.
# See https://golang.org/cmd/go/#hdr-Environment_variables
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN make build

RUN pwd
RUN ls -la

# Deploy Stage
FROM alpine:3.14 AS DeployStage
LABEL stage=deploy

# Create a non-root group
# RUN addgroup -S nonroot

# Create a non-root user and add them to the nonroot group
# RUN adduser -S -G nonroot nonroot

WORKDIR /app

# Copy the compiled Go Echo API to the working directory
COPY --from=BuildStage /app/ecnl ./ecnl

# Copy the custom configuration file to the working directory
COPY config/prod-config.yaml ./config.yaml

# Start the application as the nonroot user
#USER nonroot:nonroot

# Expose the port that your Go Echo API will listen on
EXPOSE 8080

# Define the command to run your application
ENTRYPOINT ["./ecnl", "api"]
