generate-proto:
	@buf generate
	@find . -type f -name "*.pb.go" -exec protoc-go-inject-tag -input={} -remove_tag_comment \;


# Installing Tools For Code Generation; Using Golang
setup-golang-tools:
	@echo "Setup golang tools"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest