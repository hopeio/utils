docker build -t jybl/godlv --build-arg BUILD_IMAGE=golang:1.22.1-alpine3.19 --build-arg RUN_IMAGE=alpine:3.19 -f Dockerfile .