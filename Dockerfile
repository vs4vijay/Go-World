# Builder Stage
FROM golang:1.12-alpine as builder

WORKDIR /go-world

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags "-w -s" -a -o ./bin/app ./main.go


# Deploy Stage
FROM alpine

RUN apk update && apk add --no-cache bash

COPY --from=builder /go-world/bin/app /app/

EXPOSE 9999

CMD ["/app/app"]
