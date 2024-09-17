FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy omni binary and rename to /app
COPY omni /app

# Mount config directory at /omni
VOLUME ["/omni"]

# Set working directory to /omni, so it automatically reads config from here.
WORKDIR /omni

ENTRYPOINT ["/app"]
