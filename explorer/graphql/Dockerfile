FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy explorer graphql binary and rename to /app
COPY graphql /app

EXPOSE 8080

# Mount config directory at /explorer_graphql
VOLUME ["/explorer_graphql"]

# Set working directory to /explorer_graphql, so it automatically reads config from here.
WORKDIR /explorer_graphql

ENTRYPOINT ["/app"]
