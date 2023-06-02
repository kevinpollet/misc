FROM golang:1.20-alpine3.17 as builder
WORKDIR /go/src/github.com/kevinpollet/100-continue
COPY go.mod ./
RUN go mod download
COPY main.go main.go
RUN go build -o main .

FROM alpine:3.17
COPY --from=builder /go/src/github.com/kevinpollet/100-continue/main /usr/local/bin
EXPOSE 8080
ENTRYPOINT [ "main" ]
