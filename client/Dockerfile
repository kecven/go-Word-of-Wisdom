FROM golang:alpine
LABEL authors="atetka"

WORKDIR /app

COPY . .

RUN go build -o pow-client ./src/main.go

CMD ["./pow-client"]
