VENDOR_DIR=vendor
GRPC_GATEWAY_REPO=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOC_OPTION=-I. -I$(VENDOR_DIR) -I$(VENDOR_DIR)/$(GRPC_GATEWAY_REPO)

.PHONY: install-glide
install-glide:
	go get github.com/Masterminds/glide

.PHONY: install-dep
install-dep:
	glide install

.PHONY: install-commands
install-commands:
	go install ./vendor/github.com/golang/protobuf/protoc-gen-go
	go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

proto/go:
	rm -rf gen/go && mkdir -p gen/go
	protoc -I/usr/local/include -I. \
  		-I$(GOPATH)/src \
  		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  		--go_out=plugins=grpc:gen/go \
  		echo/echo.proto

proto/gateway:
	#rm -rf gen/go && mkdir -p gen/go
	protoc -I/usr/local/include -I. \
		-I$(GOPATH)/src \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:gen/go \
		echo/echo.proto
