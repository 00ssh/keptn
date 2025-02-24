# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16.2-alpine as builder-base

WORKDIR /go/src/github.com/keptn/keptn/mongodb-datastore

# Force the go compiler to use modules
ENV GO111MODULE=on
ENV BUILDFLAGS=""
ENV GOPROXY=https://proxy.golang.org

RUN apk add --no-cache gcc libc-dev git

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy local code to the container image.
COPY . .

FROM golang:1.16.2 as builder-test

# install additional dependencies
RUN apt-get install -y gcc libc-dev git

# Force the go compiler to use modules
ENV GO111MODULE=on
ENV BUILDFLAGS=""
ENV GOPROXY=https://proxy.golang.org

# Pre-Req for gin-gonic/go-swagger: Install swag cli to generate swagger.yaml and swagger.json
RUN go get -u github.com/swaggo/swag/cmd/swag@v1.7.0

WORKDIR /go/src/github.com/keptn/keptn/shipyard-controller

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy local code to the container image.
COPY . .

CMD go test -race -coverprofile=coverage.txt -covermode=atomic -v ./... && mv ./coverage.txt /shared/coverage.txt

FROM builder-base as builder
ARG version=develop
ARG debugBuild

# set buildflags for debug build
RUN if [ ! -z "$debugBuild" ]; then export BUILDFLAGS='-gcflags "all=-N -l"'; fi

RUN sed -i "s/version: develop/version: ${version}/g" /go/src/github.com/keptn/keptn/mongodb-datastore/swagger.yaml

# Build the command inside the container.
# (You may fetch or manage dependencies here, either manually or with a tool like "godep".)
RUN cd cmd/mongodb-datastore-server && GOOS=linux go build -ldflags '-linkmode=external' $BUILDFLAGS -v -o mongodb-datastore

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.15 as production
ARG version=develop
LABEL org.opencontainers.image.source = "https://github.com/keptn/keptn" \
    org.opencontainers.image.url = "https://keptn.sh" \
    org.opencontainers.image.title="Keptn MongoDB Datastore" \
    org.opencontainers.image.vendor="Keptn" \
    org.opencontainers.image.documentation="https://keptn.sh/docs/" \
    org.opencontainers.image.licenses="Apache-2.0" \
    org.opencontainers.image.version="${version}"

# we need to install ca-certificates and libc6-compat for go programs to work properly
RUN apk add --no-cache ca-certificates libc6-compat

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/keptn/keptn/mongodb-datastore/cmd/mongodb-datastore-server/mongodb-datastore /mongodb-datastore
COPY --from=builder /go/src/github.com/keptn/keptn/mongodb-datastore/swagger-ui /swagger-ui
COPY --from=builder /go/src/github.com/keptn/keptn/mongodb-datastore/swagger.yaml /swagger-ui/swagger-original.yaml

COPY --from=builder /go/src/github.com/keptn/keptn/mongodb-datastore/swagger.yaml /swagger-ui/swagger.yaml
# Replace contents for api proxy
RUN sed -i "s|basePath: /|basePath: /api/mongodb-datastore |g" /swagger-ui/swagger.yaml
RUN sed -i '/paths:/i securityDefinitions:\n  key:\n    type: apiKey\n    in: header\n    name: x-token\nsecurity:\n  - key: []\n' /swagger-ui/swagger.yaml

EXPOSE 8080

# required for external tools to detect this as a go binary
ENV GOTRACEBACK=all

# KEEP THE FOLLOWING LINES COMMENTED OUT!!! (they will be included within the travis-ci build)
#travis-uncomment ADD MANIFEST /
#travis-uncomment COPY entrypoint.sh /
#travis-uncomment ENTRYPOINT ["/entrypoint.sh"]


RUN adduser -D nonroot -u 65532
USER nonroot

# Run the web service on container startup.
CMD ["/mongodb-datastore", "--port=8080", "--host=0.0.0.0"]
