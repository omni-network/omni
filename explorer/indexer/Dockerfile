FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy explorer indexer binary and rename to /app
COPY indexer /app

# Mount config directory at /explorer_indexer
VOLUME ["/explorer_indexer"]

# Set working directory to /explorer_indexer, so it automatically reads config from here.
WORKDIR /explorer_indexer

ENTRYPOINT ["/app"]
