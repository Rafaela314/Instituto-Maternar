FROM golang:1.17-alpine

RUN mkdir /app

WORKDIR /app

COPY . /app

COPY go.mod  ./

RUN go mod download

ENV GO111MODULE=on

RUN go build -o /instituto-maternar

EXPOSE 8888 8080 8443

CMD ["/instituto-maternar"]