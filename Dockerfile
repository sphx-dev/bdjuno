FROM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache make git gcc libc-dev
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./
ARG arch=x86_64
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 7d2239e9f25e96d0d4daba982ce92367aacf0cbd95d2facb8442268f2b1cc1fc
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep f6282df732a13dec836cda1f399dd874b1e3163504dbd9607c6af915b2740479
# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a
RUN make build

FROM alpine:latest
WORKDIR /bdjuno
RUN apk update
RUN apk add postgresql
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
COPY database/schema /var/lib/postgresql/schema

CMD [ "bdjuno", "start" ]
