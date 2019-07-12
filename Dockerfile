FROM golang:1.12-alpine3.9

COPY . /go/src/github.com/hitman99/hashing
WORKDIR /go/src/github.com/hitman99/hashing
RUN go build -o hashing

FROM alpine:3.9
LABEL maintainer="tomas@adomavicius.com"

RUN apk --no-cache add ca-certificates
WORKDIR /hashing
COPY --from=0 /go/src/github.com/hitman99/hashing/hashing hashing
COPY gen/http/openapi.json /hashing/static/openapi.json
ENV PATH="/hashing/:${PATH}"

EXPOSE 8080
CMD ["hashing", "server"]