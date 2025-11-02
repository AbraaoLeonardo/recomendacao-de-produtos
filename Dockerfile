FROM golang:tip-alpine3.22  AS builder
WORKDIR /APP
COPY . /APP
RUN go build -o bin/main

FROM alpine:3.22.2
WORKDIR /APP
COPY --from=builder /APP/bin/main .
COPY --from=builder /APP/config/ /APP/config/
EXPOSE 8080
ENTRYPOINT ["./main"] 
