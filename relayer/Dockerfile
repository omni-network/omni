FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy relayer binary and rename to /app
COPY relayer /app

# Mount config directory at /relayer
VOLUME ["/relayer"]

# Set working directory to /relayer, so it automatically reads relayer.toml from here.
WORKDIR /relayer

ENTRYPOINT ["/app"]
