# Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY . .

# テストの実行
# RUN go test -v ./...

# アプリケーションのビルド
RUN go build -o main main.go

CMD ["./main"]
