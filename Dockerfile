# Builder Stage
FROM golang:1.13 as builder

WORKDIR /go/src/go-world

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/go-world

# Deploy Stage
FROM alpine:latest

RUN apk update
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/bin/go-world /go-world

EXPOSE 9999

CMD ["/go-world"]
