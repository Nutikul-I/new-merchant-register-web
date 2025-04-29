FROM golang:1.24-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM golang:1.24-alpine

WORKDIR /app/

COPY --from=builder /go/src/app/main .
COPY --from=builder /go/src/app/template ./template

RUN chmod +x main

EXPOSE 8080

ENTRYPOINT ["/app/main"]
