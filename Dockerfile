FROM golang:1.14.4-stretch as builder

WORKDIR /github.com/momotaro98/mixlunch-email-notification

ENV GO111MODULE="on"
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o mixlunch-email-notification

FROM alpine:3.8

WORKDIR /root/

# Need CA certificates to use HTTPS in a container
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Multi stage build function of Docker
COPY --from=builder /github.com/momotaro98/mixlunch-email-notification/mixlunch-email-notification .
COPY --from=builder /github.com/momotaro98/mixlunch-email-notification/templates/body.html.tpl ./templates/

ENV TARGET_DATE=${TARGET_DATE}

CMD ./mixlunch-email-notification -target-date $TARGET_DATE
