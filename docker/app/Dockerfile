# Builder Stage
FROM golang:1.12-alpine as builder

WORKDIR /Go-World

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags "-w -s" -a -o ./bin/app ./cmd/app


# Deploy Stage
FROM alpine

RUN apk update && apk add --no-cache bash

COPY --from=builder /Go-World/bin/app /app/

EXPOSE 10000

CMD ["/Go-World/bin/app"]