FROM golang:1.11

USER nobody

RUN mkdir -p /go/src/github.com/Jairuleizer/golang-ex
WORKDIR /go/src/github.com/Jairuleizer/golang-ex

COPY . /go/src/github.com/Jairuleizer/golang-ex
RUN go build

CMD ["./golang-ex"]
