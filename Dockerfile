# Start from golang base image
FROM golang:1.17-alpine

# Install git for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app

WORKDIR /app

COPY . /app

COPY ./.env  .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

ENV GO111MODULE=on

# Build the Go app
RUN go build -o /instituto-maternar

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["/instituto-maternar"]