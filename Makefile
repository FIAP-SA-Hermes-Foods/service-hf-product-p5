build-proto:
	protoc \
	--go_out=product_api_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=product_api_proto \
	--go-grpc_opt=paths=source_relative \
	product_api.proto
