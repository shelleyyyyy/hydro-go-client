FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .

RUN go build .

CMD ["client"]
