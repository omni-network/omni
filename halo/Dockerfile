FROM scratch

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Create /tmp directory (default cometBFT --temp-dir)
COPY --from=alpine:latest /tmp /tmp

COPY halo /app

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s CMD ["/app", "ready"]

# Mount home directory at /halo
VOLUME ["/halo"]

ENTRYPOINT ["/app"]
CMD ["run"]
