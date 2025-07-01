FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o abtest-platform ./cmd/service/main.go

CMD ["./abtest-platform"]