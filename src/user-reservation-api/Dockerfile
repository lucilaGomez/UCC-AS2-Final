FROM golang:1.20

ADD ./ /api

WORKDIR /api
RUN go mod init user-reservation-api
RUN go mod tidy
RUN go build -o user-reservation-api .
RUN go test ./...
RUN chmod +x /api

ENTRYPOINT ["/api/user-reservation-api"]
