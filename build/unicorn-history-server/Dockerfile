ARG ALPINE_VERSION=3.20

FROM alpine:${ALPINE_VERSION}

WORKDIR /app

COPY assets ./assets
COPY bin/app/unicorn-history-server ./unicorn-history-server
COPY migrations ./migrations
COPY config/unicorn-history-server/config.yml ./config.yml


ENTRYPOINT ["/app/unicorn-history-server"]
CMD ["--config", "config.yml"]
