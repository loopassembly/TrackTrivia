# Use Go 1.22 as the base image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Install dependencies including packages needed for cgo
RUN apt-get update && \
    apt-get install -y \
    git \
    sqlite3 \
    libsqlite3-dev \
    build-essential

# Download Go dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Set environment variables
ENV GO_ENV=production
ENV PORT=8080
ENV CGO_ENABLED=1

# Check if files are copied correctly
RUN ls -l

# Print Go version and environment variables
RUN go version
RUN env

# Build the Go app
RUN go build -o tracktrivia main.go

# Run the Go app
CMD ["./tracktrivia"]
