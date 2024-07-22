# Start from the official Go image to create a build artifact.
FROM golang:1.22.1 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go mod and sum files to download dependencies.
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the Go app as a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start from scratch (or alpine) for a smaller final image.
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# COPY config yaml file
COPY config/config.yaml config/config.yaml

# COPY rbac_model.conf
COPY config/authorization/rbac_model.conf config/authorization/rbac_model.conf


LABEL Name=routehubservice Version=0.0.2
# Expose port 8080 to the outside world.
EXPOSE 8081

# Command to run the executable.
CMD ["./main"]

