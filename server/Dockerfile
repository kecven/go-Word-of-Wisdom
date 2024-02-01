FROM golang:alpine
LABEL authors="atetka"

WORKDIR /app

COPY . .

RUN go build -o wisdom-server ./src/main.go

CMD ["./wisdom-server"]
