FROM golang:1.21-bullseye
WORKDIR /temp-go
COPY ./tools .
RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install sudo build-essential make bash git openssh-client curl neovim unzip -y 

RUN unzip -o protoc-25.2-linux-x86_64.zip -d /usr/local bin/protoc \ 
  && unzip -o protoc-25.2-linux-x86_64.zip -d /usr/local 'include/*'
WORKDIR /workspace

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0


CMD [ "sleep infinity" ]