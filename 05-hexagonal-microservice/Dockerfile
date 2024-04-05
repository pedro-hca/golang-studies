# Use the official Go image as the base image
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and their dependencies
COPY go.mod ./
COPY go.sum ./

# Run go mod tidy to update dependencies
RUN go mod tidy

# Download Go dependencies
RUN go mod download

# Copy the rest of the application files into the container
COPY . .

# Compile the application into a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a scratch Docker image for a smaller and more secure image
FROM scratch

# Copy the compiled binary into the root directory
COPY --from=builder /app/main .

# Command to execute the application
CMD ["./main"]
