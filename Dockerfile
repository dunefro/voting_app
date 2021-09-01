FROM golang:1.14-alpine as build
WORKDIR /src
COPY server.go .
RUN go build -o /out/voting_app .
FROM alpine as bin
COPY --from=build /out/voting_app /usr/local/bin/
ENTRYPOINT /usr/local/bin/voting_app
