# Build stage 1
FROM golang:1.22-alpine3.19 as builder
LABEL stage=builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN apk add --update gcc musl-dev make

WORKDIR /app
COPY go.* /app/
COPY cmd /app/cmd
COPY internal/ /app/internal
#COPY docs /app/docs

COPY Makefile /app/Makefile

RUN make build

# Build stage 2
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/

WORKDIR /app
COPY .env .env
COPY db /app/db
COPY configs /app/configs

CMD ["./main"]