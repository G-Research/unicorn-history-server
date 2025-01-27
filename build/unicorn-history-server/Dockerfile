ARG NODE_VERSION=22
ARG ALPINE_VERSION=3.20

FROM node:${NODE_VERSION} AS uhs-web

COPY web /build/src

WORKDIR /build/src

RUN npm install -g pnpm            && \
    pnpm install                   && \
    pnpm run build:prod

# Install yunikorn-web
RUN git clone https://github.com/G-Research/yunikorn-web.git /build/src/yunikorn-web

RUN cd /build/src/yunikorn-web && \
    pnpm install && \
    pnpm run build:prod

FROM alpine:${ALPINE_VERSION}

COPY --from=uhs-web /build/assets /app/assets
COPY --from=uhs-web /build/src/yunikorn-web/dist/yunikorn-web /app/assets
COPY bin/app/unicorn-history-server /app/unicorn-history-server
COPY migrations /app/migrations
COPY config/unicorn-history-server/config.yml /app/config.yml

WORKDIR /app


ENTRYPOINT ["/app/unicorn-history-server"]
CMD ["--config", "config.yml"]
