FROM golang:1.18.2-alpine3.15 AS build

WORKDIR /generate

COPY go.mod ./
ADD . .
RUN go mod download && go mod verify
RUN go build -v -o generate

FROM alpine:3.15.4
COPY --from=build /generate /usr/local/bin
ENTRYPOINT ["generate"]
