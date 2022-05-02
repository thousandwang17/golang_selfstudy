FROM golang:1.17.8-alpine

COPY ./ /src/
WORKDIR /src/cmd/client

RUN CGO_ENABLED=0 GOOS=linux go build -o kitClient

FROM alpine:3.14
COPY --from=0 /src/cmd/client/kitClient /
CMD ./kitClient