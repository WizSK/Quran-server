# syntax=docker/dockerfile:1

FROM golang:1.20.4

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./quran-server

EXPOSE 8001

CMD [ "./quran-server"]
