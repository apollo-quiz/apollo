#!/bin/bash

protoc --plugin=../ui/quiz-admin/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=../ui/quiz-admin/src/protobuf --proto_path=../data/proto ../data/proto/*.proto
protoc --plugin=../ui/quiz-participant/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=../ui/quiz-participant/src/protobuf --proto_path=../data/proto ../data/proto/*.proto