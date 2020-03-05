FROM golang:1.14-alpine AS build

WORKDIR /go/src/github.com/Azaradel/mm-slash-command

RUN adduser -D -g '' app
RUN apk add --no-cache tzdata git ca-certificates && update-ca-certificates

COPY . .

RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/slash-command

FROM scratch

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /go/bin/slash-command /go/bin/slash-command

USER app
EXPOSE 7890
ENTRYPOINT [ "/go/bin/slash-command" ]