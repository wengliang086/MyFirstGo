# GOGO_ROOT=${GOPATH}/src/github.com/gogo/protobuf
GOGO_ROOT=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.2.1
# /Users/phoenix/go/pkg/mod/github.com/gogo/protobuf@v1.2.1

protoc -I.:${GOGO_ROOT} --gogofaster_out=plugins=grpc:. helloworld.proto
