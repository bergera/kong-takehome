# build image
FROM golang:1.17-alpine3.15 as builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./kong_takehome

# runtime image
FROM alpine:3.15
EXPOSE 8080
RUN addgroup kong_takehome
RUN adduser -D -G kong_takehome kong_takehome
COPY --from=builder --chown=kong_takehome:kong_takehome /usr/local/bin/app /usr/local/bin/app
USER kong_takehome
CMD ["/usr/local/bin/app"]
