FROM golang:1.12-alpine

WORKDIR /Go-World

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY . .

RUN go build -ldflags "-w -s" -a -o ./bin/app ./cmd/app

CMD ["/Go-World/bin/app"]

EXPOSE 10000