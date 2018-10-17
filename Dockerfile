FROM golang:alpine as builder

ADD . /go/src/CompanyService
WORKDIR /go/src/CompanyService

# build the source
RUN CGO_ENABLED=0 GOOS=linux go install -a -installsuffix cgo ./...

# use a minimal alpine image
FROM alpine

# add ca-certificates in case you need them
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# set working directory
WORKDIR /go/bin

USER 1001
# run the binary
CMD ["./companyservice"]

EXPOSE 8080