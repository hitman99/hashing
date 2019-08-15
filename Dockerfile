FROM golang:1.12-alpine3.9

COPY . /hashing
WORKDIR /hashing
RUN go build -mod=vendor -o hashing

FROM alpine:3.9
LABEL maintainer="tomas@adomavicius.com"

RUN apk --no-cache add ca-certificates
WORKDIR /hashing
COPY --from=0 /hashing/hashing hashing
COPY gen/http/openapi.json /hashing/static/openapi.json
ENV PATH="/hashing/:${PATH}"

EXPOSE 8080
CMD ["hashing", "server"]