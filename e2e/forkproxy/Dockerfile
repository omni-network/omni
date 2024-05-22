FROM --platform=linux/amd64 ghcr.io/foundry-rs/foundry:latest

# Copy forkproxy binary and rename to /app
COPY forkproxy /app

EXPOSE 8545

# Mount config directory at /forkproxy
VOLUME ["/forkproxy"]

# Set working directory to /forkproxy, so it automatically reads config from here.
WORKDIR /forkproxy

ENTRYPOINT ["/app"]
