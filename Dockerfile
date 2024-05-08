FROM golang:1.21.3-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/alishchenko/discountaria
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/discountaria /go/src/github.com/alishchenko/discountaria


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/discountaria /usr/local/bin/discountaria
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["discountaria"]
