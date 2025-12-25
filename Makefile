PROTO_FILE = auth.proto
PROTO_PATH = ./internal/transport/grpc/proto/${PROTO_FILE}
PROTO_OUT = ./internal/transport/grpc/pb

build_protoc:
	rm -f ${PROTO_OUT}/user_grpc.pb.go
	rm -f ${PROTO_OUT}/user.pb.go
	protoc --go_out=${PROTO_OUT} --go_opt=paths=import --go-grpc_out=${PROTO_OUT} --go-grpc_opt=paths=import $(PROTO_PATH)