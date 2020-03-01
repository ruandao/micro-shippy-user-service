FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO11MODULE=ON

COPY .. .

RUN go mod download
RUN cd ser
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-user

# RUN container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-user .

CMD ["./shippy-service-user"]
