run: build
	@./bin/svc_1

build:
	go build -o bin/svc_1 .

clean:
	rm -f bin/svc_1

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service_proto.proto

.PHONY: proto


