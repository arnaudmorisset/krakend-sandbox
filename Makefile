.PHONY: proto

proto:
	@protoc --go_out=./greeter/ --go_opt=paths=source_relative \
					--go-grpc_out=./greeter/ --go-grpc_opt=paths=source_relative \
					--descriptor_set_out=./gateway/config/grpc_catalog/proto.pb \
					--include_source_info --include_imports \
					-I. proto/greeter.proto

build:
	@docker build -t greeter -f ./greeter/Dockerfile ./greeter --no-cache

build-hello-plugin:
	@docker run -it -v "${PWD}/gateway/config/plugins/hello:/app" -w /app \
		krakend/builder:2.4.2 \
		go build -buildmode=plugin -o hello-plugin.so .

start:
	@docker-compose up

grpcurl-greeter:
	@grpcurl -plaintext \
					 -import-path ./proto -proto greeter.proto \
					 -d '{"name": "Arnaud"}' \
					 localhost:50051 greeter.Greeter/SayHello

grpcurl-user:
	@grpcurl -plaintext \
					 -import-path ./proto -proto greeter.proto \
					 -d '{"id": "1"}' \
					 localhost:50051 greeter.User/GetUser

grpcurl-profile:
	@grpcurl -plaintext \
					 -import-path ./proto -proto greeter.proto \
					 -d '{"id": "1"}' \
					 localhost:50051 greeter.Profile/GetProfile

grpcurl-recurse:
	@grpcurl -plaintext \
					 -import-path ./proto -proto greeter.proto \
					 localhost:50051 greeter.Recurse/GetRecursiveData

grpcurl-generic:
	@grpcurl -plaintext \
					 -import-path ./proto -proto greeter.proto \
					 -d '{"data": {"string_value": "1"} }' \
					 localhost:50051 greeter.Generic/GetGenericData

curl-greeter:
	@curl -X GET "localhost:8080/greeting"

curl-users:
	@curl -X GET "localhost:8080/users/1"

curl-recurse:
	@curl -X GET "localhost:8080/recurse"

curl-generic:
	@curl -X GET "localhost:8080/generic" \
		-d '{"data": {"string_value": "1"} }'
