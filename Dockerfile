FROM --platform=$BUILDPLATFORM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache make git gcc libc-dev
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./
RUN go mod download
ARG arch=x86_64
# we use the same arch in the CI as a workaround since we don't use the wasm in the indexer
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.5.0/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.${arch}.a
# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a
RUN make build

FROM --platform=$TARGETPLATFORM alpine:latest
WORKDIR /bdjuno
RUN apk update
RUN apk add postgresql
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
COPY database/schema /var/lib/postgresql/schema
RUN chmod a+rx /var/lib/postgresql && \
    chmod a+rx /var/lib/postgresql/schema

CMD [ "bdjuno", "start" ]
