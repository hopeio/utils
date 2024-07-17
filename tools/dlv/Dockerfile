
ARG BUILD_IMAGE=jybl/protogen
ARG RUN_IMAGE=frolvlad/alpine-glibc
FROM ${BUILD_IMAGE} AS builder

ENV GOPROXY https://goproxy.io,https://goproxy.cn,direct

RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM ${RUN_IMAGE} AS runtime

COPY --from=builder /go/bin/dlv /bin/