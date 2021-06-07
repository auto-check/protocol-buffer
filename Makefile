.PHONY: all

TARGET = auth

all: $(TARGET)

auth:
	protoc --go_out golang/$@ proto/$@/$@.proto
	protoc --go-grpc_out golang/$@ proto/$@/$@.proto
	protoc --grpc-gateway_out golang/$@ proto/$@/$@.proto

main:
	protoc --go_out golang/$@ proto/$@/$@.proto
	protoc --go-grpc_out golang/$@ proto/$@/$@.proto
	protoc --grpc-gateway_out golang/$@ proto/$@/$@.proto
