FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOARCH=arm64 GOOS=linux go build -o main ./cmd/api/main.go
# RUN go build -o main ./cmd/api/main.go

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]