FROM golang:1.18 AS builder
WORKDIR $GOPATH/src/commentci
COPY go.mod go.sum ./
RUN go get -d -v -u all
COPY ./ ./
RUN CGO_ENABLED=0 go build -tags netgo -a -o /go/bin/commentci

FROM alpine:3.15
COPY --from=builder /go/bin/commentci /go/bin/commentci
ENTRYPOINT ["/go/bin/commentci"]
CMD ["-h"]
