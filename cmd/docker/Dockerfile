FROM golang:alpine3.12 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

COPY ../../go.mod .
COPY ../../go.sum .
RUN go mod download

COPY ../.. .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o cygo-bin .

FROM alpine:3.12 AS final

WORKDIR /app
COPY --from=builder /build/cygo-bin /app/
COPY --from=builder /build/assert /app/assert
COPY --from=builder /build/.env /app/.env
COPY --from=builder /etc/passwd /etc/passwd

USER app-runner
EXPOSE 80
ENTRYPOINT ["/app/cygo-bin"]