FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o ab_test_platform ./cmd/service/main.go

CMD ["./ab_test_platform"]