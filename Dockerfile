FROM golang:1.16-alpine

RUN apk add --update make

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

EXPOSE 8000
CMD ["bin/api"]
