FROM golang:1.20

ADD ./ /api

WORKDIR /api
RUN go mod init hotel-api
RUN go mod tidy
RUN go build -o hotel-api .
RUN go test ./...
RUN chmod +x /api

ENTRYPOINT ["/api/hotel-api"]
