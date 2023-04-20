PROTOC_FLAGS = -Iproto
GO_OUT_FLAGS = --go_out=. --go_opt=module=github.com/sudipto-003/grpc-exp/expunary
GRPC_OUT_FLAGS = --go-grpc_out=. --go-grpc_opt=module=github.com/sudipto-003/grpc-exp/expunary

.PHONY: build
build: expunary

expunary: proto/asmd.pb.go server/main.go client/main.go
	go build -o bin/unary/server ./server
	go build -o bin/unary/client ./client

proto/asmd.pb.go: proto/asmd.proto
	protoc $(PROTOC_FLAGS) $(GO_OUT_FLAGS) $(GRPC_OUT_FLAGS) $<

.PHONY: clean
clean:
	rm -f proto/asmd.pb.go proto/asmd_grpc.pb.go bin/unary/server bin/unary/client
