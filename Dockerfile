FROM golang:1.10-alpine

RUN apk add --no-cache git

WORKDIR /app
CMD ["./app"]

COPY . .

RUN go get -v -d
RUN go build -v
