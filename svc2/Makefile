run: build
	@./bin/svc_2

build:
	go build -o bin/svc_2 .

clean:
	rm -f bin/svc_2

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service_proto.proto

.PHONY: proto


