FROM --platform=linux/amd64 ghcr.io/foundry-rs/foundry:latest

# Copy anvilproxy binary and rename to /app
COPY anvilproxy /app

EXPOSE 8545

# Mount config directory at /anvilproxy
VOLUME ["/anvilproxy"]

# Set working directory to /anvilproxy, so it automatically reads config from here.
WORKDIR /anvilproxy

ENTRYPOINT ["/app"]
