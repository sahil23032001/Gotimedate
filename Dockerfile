# ----------------------------------------------------------------------
# Stage 1: The Builder Stage
# Purpose: Compile the Go application into a single static binary.
# ----------------------------------------------------------------------
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to enable caching
# Assuming these files are present in your local directory
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application source code (including main.go)
COPY . .

# Build the static binary. CGO_ENABLED=0 is essential for minimal base images.
RUN CGO_ENABLED=0 go build -o datetimeapp .

# ----------------------------------------------------------------------
# Stage 2: The Final Stage
# Purpose: Create the final, minimal image by hosting only the compiled binary.
# ----------------------------------------------------------------------
FROM alpine:latest

# The application listens on port 8080 (as defined in main.go)
EXPOSE 8080

# Set the working directory for the final container
WORKDIR /app

# Copy the compiled binary from the 'builder' stage into the final image
COPY --from=builder /app/datetimeapp .

# Define the entrypoint to run the compiled application.
ENTRYPOINT ["./datetimeapp"]
