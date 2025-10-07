FROM golang:1.25.0-alpine AS build
WORKDIR /workspace
RUN apk add --no-cache git ca-certificates && update-ca-certificates
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOTOOLCHAIN=auto
ENV CGO_ENABLED=0


COPY common/go.mod ./common/go.mod
COPY order-service/go.mod ./order-service/go.mod
RUN cd order-service && go mod download


COPY common/ ./common/
COPY order-service/ ./order-service/


RUN cd order-service && go build -trimpath -ldflags="-s -w" -o /workspace/bin/order .
ARG GOOSE_VERSION=v3.20.0
RUN go install github.com/pressly/goose/v3/cmd/goose@${GOOSE_VERSION}


FROM alpine:3.20 AS release
WORKDIR /app
RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build /workspace/bin/order /usr/local/bin/order
COPY --from=build /go/bin/goose /usr/local/bin/goose
COPY order-service/migrations ./migrations

COPY order-service/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

ENV HTTP_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]


FROM golang:1.25.0-alpine AS dev
WORKDIR /workspace
RUN apk add --no-cache git bash ca-certificates && update-ca-certificates
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOTOOLCHAIN=auto

COPY common/go.mod ./common/go.mod
COPY order-service/go.mod ./order-service/go.mod

COPY common/ ./common/
COPY order-service/ ./order-service/

ARG AIR_VERSION=v1.63.0
ARG GOOSE_VERSION=v3.20.0
RUN go install github.com/air-verse/air@${AIR_VERSION} \
 && go install github.com/pressly/goose/v3/cmd/goose@${GOOSE_VERSION}

WORKDIR /workspace/order-service
COPY order-service/entrypoint.dev.sh /usr/local/bin/entrypoint.dev.sh
RUN chmod +x /usr/local/bin/entrypoint.dev.sh

ENV HTTP_ADDR=:8080
EXPOSE 8080
CMD ["/usr/local/bin/entrypoint.dev.sh"]
