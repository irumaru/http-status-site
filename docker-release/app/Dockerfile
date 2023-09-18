FROM golang:1.21.1-alpine3.18 AS builder

WORKDIR /usr/src/build
COPY ./src /usr/src/build

RUN go build main.go

FROM alpine:3.18
WORKDIR /http
COPY --from=builder /usr/src/build/main /http/main

CMD [ "./main" ]
