FROM golang:1.12.6
ARG version
COPY go.* /go/src/github.com/flowerinthenight/golang-monorepo/
COPY pkg/ /go/src/github.com/flowerinthenight/golang-monorepo/pkg/
COPY vendor/ /go/src/github.com/flowerinthenight/golang-monorepo/vendor/
COPY cmd/samplecmd/ /go/src/github.com/flowerinthenight/golang-monorepo/cmd/samplecmd/
WORKDIR /go/src/github.com/flowerinthenight/golang-monorepo/cmd/samplecmd/
RUN GO111MODULE=on GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-X github.com/flowerinthenight/golang-monorepo/cmd/samplecmd/main.version=$version" -a -installsuffix cgo -o samplecmd .

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /samplecmd/
COPY --from=0 /go/src/github.com/flowerinthenight/golang-monorepo/cmd/samplecmd .
ENTRYPOINT ["/samplecmd/samplecmd"]
CMD ["--logtostderr"]
