#!/usr/bin/bash

### run this script from the project root path.

test=./djpm/run/testdata
npm install --prefix $test;

protoc --plugin=protoc-gen-as=$test/node_modules/.bin/as-proto-gen --proto_path=./proto/aspect/v1/ --as_out=$test/runner hostapi_message.proto
protoc --plugin=protoc-gen-as=$test/node_modules/.bin/as-proto-gen --proto_path=./proto --as_out=$test/runner scheduler/v1/schedule_message.proto

npm --prefix $test run asbuild:release