# syntax=docker/dockerfile:1

FROM golang:1.20.4-alpine3.18 AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./quran-server

FROM scratch

COPY --from=builder /build /

EXPOSE 8001

CMD [ "./quran-server"]
