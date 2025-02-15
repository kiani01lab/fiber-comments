FROM golang:1.23.2-alpine

LABEL maintainer="Ali Kiani <codewithkiani@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/server/main.go

EXPOSE 8000

CMD [ "/app/main" ]