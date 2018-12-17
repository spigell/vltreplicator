FROM golang:alpine as builder
RUN apk --no-cache add git make gcc

WORKDIR /go/vltreplicator

COPY . /go/vltreplicator

RUN make submodule_check
RUN make

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /root/
COPY --from=builder /go/vltreplicator/bin/vltreplicator .
COPY /docker/entrypoint.sh /tmp/entrypoint.sh

ENTRYPOINT ["/tmp/entrypoint.sh"]
