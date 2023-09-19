FROM golang:1.21.1-alpine3.18 AS builder
# AlpineにはDebianなどで使用されているC標準ライブラリ/lib64/ld-linux-x86-64.so.2が存在しないため
# 最初からAlpineでビルドして、互換ライブラリを使用するようにする必要がある。

WORKDIR /usr/src/build
COPY ./src /usr/src/build

RUN go build main.go

FROM alpine:3.18
WORKDIR /http
COPY --from=builder /usr/src/build/main /http/main

EXPOSE 8080

CMD [ "./main" ]
