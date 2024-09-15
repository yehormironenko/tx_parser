FROM golang:alpine AS build
ARG CI_GITLAB_MODULES_REG_USERNAME
ARG CI_GITLAB_MODULES_REG_TOKEN

ENV GOPRIVATE gitlab.com/bango

COPY ./ /src
WORKDIR /src
RUN apk update && apk add --no-cache git ca-certificates build-base

RUN git config --global \
            url."https://${CI_GITLAB_MODULES_REG_USERNAME}:${CI_GITLAB_MODULES_REG_TOKEN}@gitlab.com/".insteadOf "https://gitlab.com/" && CGO_ENABLED=1 GOOS=linux go build -a -tags musl -installsuffix cgo -ldflags '-w -s -extldflags "-static"' -o /main cmd/main.go

# Create a non-root user (in the build stage)
RUN adduser -D -g '' appuser

FROM scratch AS runtime

# Copy the built binary and other necessary files
COPY --from=build /main /
COPY --from=build /src/config/*.yaml /config/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the non-root user from the build stage
COPY --from=build /etc/passwd /etc/passwd

# Use the non-root user to run the application
USER appuser

EXPOSE 8085
CMD ["/main"]
