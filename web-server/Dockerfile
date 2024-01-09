FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./web-server .

# Build the binary.
#RUN CGO_ENABLED=0 GOOS=linux go build -o main web-server/internal/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]