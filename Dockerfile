FROM golang:1.16-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./
COPY database/schema /database/schema
RUN make build

FROM alpine:latest
WORKDIR /bdjuno
RUN apk update
RUN apk add postgresql
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
COPY --from=builder /database/schema /var/lib/postgresql/schema

CMD [ "bdjuno" ]
