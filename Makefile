.PHONY: generate-client
generate-client:
	@protoc \
  --proto_path=${contracts} \
  --go_out=${PWD}/internal/generated/investapi \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_out=${PWD}/internal/generated/investapi \
  ${contracts}*.proto
