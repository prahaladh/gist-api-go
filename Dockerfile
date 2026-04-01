# --- Stage 1: Build Stage ---
FROM golang:1.25-alpine AS builder

# Install build dependencies and update for security patches
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create a non-privileged user to run the app later
RUN adduser -D -g '' devops

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# If any test fails, the Docker build will stop here.
RUN go test -v ./...

# Build the binary:
RUN  go build \
    -ldflags="-w -s" \
    -o gist-api .

# --- Stage 2:  Run  ---
FROM scratch

# Import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Import CA certificates so the app can talk to GitHub via HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy only the compiled binary
COPY --from=builder /app/gist-api /gist-api

# Use the non-root user we created
USER devops

# Document the port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/gist-api"]