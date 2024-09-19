# # Use the official Golang image as a build stage
# FROM golang:1.23 AS builder
#
# # Set the working directory
# WORKDIR /app
#
# RUN go install github.com/air-verse/air@latest
# # Copy go mod and sum files
# COPY go.mod go.sum ./
# # Download dependencies
# RUN go mod download
# # # Copy the rest of the application
# COPY . .
#
# # Build the application
# RUN go build -o main .
# # CMD ["air"]
#
# CMD ["/app/main"]
# # #
# # # # # Use a minimal image for the final stage
# # FROM scratch
# # WORKDIR /app
# # # #
# # COPY --from=builder /app/main .
# # # #
# # # # Expose the port your app listens on
# # # EXPOSE 8080
# # #
# # # # Run the application
# # CMD ["/app/main"]
# Stage 1: Build the Go app
FROM golang:1.23-alpine AS builder

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Stage 2: Run the app with live reload
FROM golang:1.23-alpine

# Install Air
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire application code
COPY --from=builder /app .

# Expose the service port
EXPOSE 8080

# Command to run the Go application with Air live reload
CMD ["air"]
