FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go get .

RUN go build -o test_api

EXPOSE 8080

CMD ./test_api