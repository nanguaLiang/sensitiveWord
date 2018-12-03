FROM golang:1.10.0
ADD . /app
RUN go get github.com/golang/protobuf/proto
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
ADD . /go/src/Exercise/RPC/sensitiveWord
RUN go install Exercise/RPC/sensitiveWord

ENTRYPOINT ["/go/bin/sensitiveWord"]