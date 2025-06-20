FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy trade binary and rename to /app
COPY trade /app

# Mount config directory at /trade
VOLUME ["/trade"]

# Set working directory to /trade, so it automatically reads config from here.
WORKDIR /trade

ENTRYPOINT ["/app"]
