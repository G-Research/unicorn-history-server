ARG ALPINE_VERSION=3.20

FROM alpine:${ALPINE_VERSION}

COPY assets /app/assets
COPY bin/app/unicorn-history-server /app/unicorn-history-server
COPY migrations /app/migrations
COPY config/unicorn-history-server/config.yml /app/config.yml

WORKDIR /app

ENTRYPOINT ["/app/unicorn-history-server"]
CMD ["--config", "config.yml"]
