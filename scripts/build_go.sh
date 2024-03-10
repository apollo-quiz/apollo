#!/bin/bash

rm -rf ../services/common/answers
rm -rf ../services/common/questions
protoc --go_out=../services --proto_path=../data/proto ../data/proto/*.proto