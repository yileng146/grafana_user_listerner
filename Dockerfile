FROM golang:1.19.5-alpine3.17 AS build-env
COPY /src /go/src/graf_user_listerner
# RUN apk add gcc musl-dev 
WORKDIR /go/src/graf_user_listerner
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o graf_user_listerner

FROM alpine:3.17.1
MAINTAINER MIKAKO
COPY --from=build-env /go/src/graf_user_listerner/graf_user_listerner /root/
WORKDIR /root/
CMD sleep 3600s
