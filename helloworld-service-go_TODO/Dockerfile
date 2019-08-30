
FROM golang:1.12-stretch AS gobuilder

WORKDIR /go/src/github.com/bygui86/go-metrics
COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/app .

# ---

FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache bash
RUN apk add --no-cache curl
# RUN apk add --no-cache ca-certificates

WORKDIR /usr/bin/
COPY --from=gobuilder /bin/app .
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# kubernetes probes
EXPOSE 7090
# monitoring
EXPOSE 7091
# rest
EXPOSE 7001

USER 1001

ENTRYPOINT "/usr/bin/app"
