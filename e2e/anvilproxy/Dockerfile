ARG FOUNDRY_VERSION
FROM --platform=linux/amd64 docker.io/chrmllr/foundry-${FOUNDRY_VERSION}
# FROM ghcr.io/foundry-rs/foundry:${FOUNDRY_VERSION}

# Need wget for localhost perturbations
# apk for apline (https://github.com/foundry-rs/foundry/blob/master/Dockerfile)
# apt for ubuntu (https://github.com/foundry-rs/foundry/blob/master/Dockerfile.cross) - apt update runs in parent
USER root
RUN if command -v apk > /dev/null; then apk add --no-cache wget; \
    elif command -v apt > /dev/null; then apt install -y wget; \
    fi
USER foundry

# Copy anvilproxy binary and rename to /app
COPY anvilproxy /app

EXPOSE 8545

# Mount config directory at /anvilproxy
VOLUME ["/anvilproxy"]

# Set working directory to /anvilproxy, so it automatically reads config from here.
WORKDIR /anvilproxy

ENTRYPOINT ["/app"]
