FROM golang:1.8

RUN \
  mkdir -p /build && \
  mkdir -p /go/src/github.com/uswitch && \
  ln -s /go/src/app /go/src/github.com/uswitch/journald-forwarder

WORKDIR /go/src/app
COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install
