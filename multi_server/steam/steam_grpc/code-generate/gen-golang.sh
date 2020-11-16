#!/usr/bin/env bash

protoDir="../proto"
outDir="../proto"
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}