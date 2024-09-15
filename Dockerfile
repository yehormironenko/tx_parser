# Build stage
FROM golang:alpine AS build

# Install dependencies
RUN apk update && apk add --no-cache git ca-certificates build-base

WORKDIR /src
COPY ./ /src

# Build the Go application with flags for static linking and optimization
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o /main ./cmd/main.go

# Runtime stage
FROM alpine:latest AS runtime

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the built binary and configuration files
COPY --from=build /main /
COPY --from=build /src/config/*.json /config/

EXPOSE 8092

CMD ["/main"]