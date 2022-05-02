FROM golang:1.17.8-alpine

COPY ./ /src/
WORKDIR /src/cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -o kitServer

FROM alpine:3.14
COPY --from=0 /src/cmd/kitServer /
CMD ./kitServer