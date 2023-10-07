# Build Stage
FROM golang:1.21.2 AS BuildStage
LABEL stage=build

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN make deps
RUN make swagger
RUN make build

# Deploy Stage
FROM alpine:3.14 AS DeployStage
LABEL stage=deploy

WORKDIR /app

COPY --from=BuildStage /app/ecnl ./ecnl
COPY config/prod-config.yaml ./config.yaml

USER nonroot:nonroot

# Expose the port that your Go Echo API will listen on
EXPOSE 8080

# Define the command to run your application
ENTRYPOINT ["./ecnl", "api"]
