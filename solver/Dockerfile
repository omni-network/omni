FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy solver binary and rename to /app
COPY solver /app

# Mount config directory at /solver
VOLUME ["/solver"]

# Set working directory to /solver, so it automatically reads config from here.
WORKDIR /solver

ENTRYPOINT ["/app"]
