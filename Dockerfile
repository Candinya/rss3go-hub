FROM golang:1.16-alpine AS BUILDER

# Set the Current Working Directory inside the container
WORKDIR /rss3go_hub

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Install basic packages
RUN apk add \
    gcc \
    g++

# Download all the dependencies
RUN go get

# Build image
RUN go build .

FROM alpine:latest AS RUNNER

COPY --from=builder /rss3go_hub/rss3go_hub .
COPY config.yml .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./rss3go_hub"]