# Stage 1: Build
FROM golang:1.23-alpine AS build

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# Install necessary tools
RUN apk add --no-cache git
# Create and switch to a new working directory
WORKDIR /app
# Copy Go module files and download dependencies
COPY go.mod go.sum main.go config.yaml utils models converter cmd ./
RUN go mod download
# Copy the source code
COPY . .
# Build the application
RUN go build -ldflags="-s -w" -o /app/GoFromMediumToHugo ./main.go

# Stage 2: Run
FROM gcr.io/distroless/static:nonroot
# Copy the binary from the build stage
COPY --from=build /app/GoFromMediumToHugo /GoFromMediumToHugo
COPY --from=build /app/config.yaml /config.yaml
# Expose a port (optional, if your app uses networking)
EXPOSE 8080
# Run the binary
ENTRYPOINT ["/GoFromMediumToHugo"]