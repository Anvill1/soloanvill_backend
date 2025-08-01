# Stage 1: Build the binary
FROM golang:1.24.4 AS builder

LABEL stage="gobuilder" org.opencontainers.image.authors=rtav3d@gmail.com

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/soloanvill_backend ./cmd/main.go 

# Stage 2: Run the app

FROM alpine:3.22

COPY --from=builder /app/soloanvill_backend /usr/local/bin/

RUN apk update && apk update --no-cache && apk add --no-cache tzdata && \
    adduser -D -u 1001 -G root soloanvill && \
    mkdir -p /app && \
    mkdir /etc/soloanvill && \
    chmod +x /usr/local/bin/soloanvill_backend && \
    chown -R soloanvill:0 /app && \
    chown -R soloanvill:0 /etc/soloanvill

ENV TZ=Europe/Moscow
ENV SOLOANVILL_DATABASE_NAME=soloanvill

WORKDIR /etc/soloanvill

EXPOSE 8080

USER 1001

CMD ["/usr/local/bin/soloanvill_backend"]
