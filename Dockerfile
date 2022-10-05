FROM golang:1.19.2-alpine3.15 AS build

WORKDIR /generate

COPY go.mod ./
ADD . .
RUN go mod download && go mod verify
RUN go build -v -o generate

FROM alpine:3.16.2
COPY --from=build /generate /usr/local/bin
ENTRYPOINT ["generate"]
