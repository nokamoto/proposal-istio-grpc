#!/bin/bash

set -eux

prototool format -d protobuf/*.proto || prototool format -w protobuf/*.proto

protoc --go_out=plugins=grpc:protobuf -I protobuf protobuf/*.proto

dep ensure

for PACKAGE in {server,client}
do
    gofmt -d ${PACKAGE}
    gofmt -w ${PACKAGE}
    golint -set_exit_status ${PACKAGE}
    go build ./${PACKAGE}/*.go
done

docker-compose build
