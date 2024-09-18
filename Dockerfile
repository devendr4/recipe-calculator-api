# Use the official Golang image as a build stage
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

RUN go install github.com/air-verse/air@latest
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# # Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .
# CMD ["air"]

CMD ["/app/main"]
# #
# # # # Use a minimal image for the final stage
# FROM scratch
# WORKDIR /app
# # #
# COPY --from=builder /app/main .
# # #
# # # Expose the port your app listens on
# # EXPOSE 8080
# #
# # # Run the application
# CMD ["/app/main"]
