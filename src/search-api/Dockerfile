FROM golang:1.20

ADD ./ /api

WORKDIR /api
RUN go mod init search-api
RUN go mod tidy
RUN go build -o search-api .
RUN go test ./...
RUN chmod +x /api

ENTRYPOINT ["/api/search-api"]
