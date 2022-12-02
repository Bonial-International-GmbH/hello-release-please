FROM golang:1.19.3-alpine3.15 as builder

WORKDIR /src

RUN apk --update --no-cache add git make

ENV CGO_ENABLED=0

COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

RUN go mod download

COPY *.go ./

RUN make build

FROM alpine:3.17.0

COPY --from=builder /src/hello-release-please /hello-release-please

ENTRYPOINT ["/hello-release-please"]
