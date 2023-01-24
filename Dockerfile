# Start from golang base image
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build the Go app
RUN go build -o /instituto-maternar

EXPOSE 8080

CMD ["/instituto-maternar"]





