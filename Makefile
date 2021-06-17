TARGET = main auth

.PHONY: all
all: $(TARGET)

.PHONY: main
main:
	protoc --go_out golang/$@ proto/$@/$@.proto
	protoc --go-grpc_out golang/$@ proto/$@/$@.proto
	protoc --grpc-gateway_out golang/$@ proto/$@/$@.proto

.PHONY: auth
auth:
	protoc --go_out golang/$@ proto/$@/$@.proto
	protoc --go-grpc_out golang/$@ proto/$@/$@.proto
	protoc --grpc-gateway_out golang/$@ proto/$@/$@.proto
