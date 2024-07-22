SHELL:=/bin/sh

.SILENT:

all: gen-proto

gen-proto:
	echo "\033[0;34m*** Gen Proto***\033[0;0m"; \
	protoc --proto_path=./util/protoc/proto/ --go_out=./util/protoc/pb/ --go_opt=paths=source_relative --go-grpc_out=./util/protoc/pb/ --go-grpc_opt=paths=source_relative ./util/protoc/proto/*
	echo "\033[0;32mFINISHED\033[0;0m"