# Start from the golang alpine base image for the builder stage
FROM golang:1.22-alpine AS builder

# Add Maintainer Info
LABEL maintainer="syahidmirdan@gmail.com"

# Set the Current Working Directory inside the builder container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the builder container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:3.8

WORKDIR /root/

# Install software without updating the package list, which could lead to vulnerable packages being installed
RUN apk add --no-cache \
    bash \
    curl \
    openssh-client \
    git=2.18.1-r0 \
    ca-certificates \
    && rm -rf /var/cache/apk/*

# Copy the prebuilt binary from the builder stage
COPY --from=builder /app/main .

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["./main"]