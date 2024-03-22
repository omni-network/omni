FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy monitor binary and rename to /app
COPY monitor /app

# Mount config directory at /monitor
VOLUME ["/monitor"]

# Set working directory to /monitor, so it automatically reads config from here.
WORKDIR /monitor

ENTRYPOINT ["/app"]
