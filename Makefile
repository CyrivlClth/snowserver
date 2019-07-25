GRPC_PROTO_PATH=./grpc/pb
PROTO_FILE=$(GRPC_PROTO_PATH)/snow.proto
.PHONY: gen-proto
go-proto:
	protoc --go_out=plugins=grpc:. $(PROTO_FILE)