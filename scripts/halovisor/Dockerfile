# Docker build args
ARG OMNI_COSMOVISOR_VERSION=v0.4.0
ARG HALO_VERSION_0_GENESIS=v0.8.1
ARG HALO_VERSION_1_ULUWATU=main

# Build stages
FROM omniops/cosmovisor:${OMNI_COSMOVISOR_VERSION} AS build-cosmovisor
FROM omniops/halo:${HALO_VERSION_0_GENESIS} AS build-0-genesis
FROM omniops/halo:${HALO_VERSION_1_ULUWATU} AS build-1-uluwatu

# Runtime stage
FROM scratch AS runtime

# Install ca-certificates (for https to rollups)
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Create /tmp directory (default cometBFT --temp-dir)
COPY --from=alpine:latest /tmp /tmp

# Cosmovisor environment variables
ENV DAEMON_HOME=/halo
ENV DAEMON_NAME=halo
ENV DAEMON_ALLOW_DOWNLOAD_BINARIES=false
ENV DAEMON_RESTART_AFTER_UPGRADE=true
ENV UNSAFE_SKIP_BACKUP=true
ENV COSMOVISOR_TIMEFORMAT_LOGS=rfc3339nano
ENV COSMOVISOR_CUSTOM_ROOT=/halovisor
ENV COSMOVISOR_CUSTOM_CURRENT_LINK=/halo/halovisor-current
ENV COSMOVISOR_CUSTOM_CURRENT_UPGRADEINFO=/halo/halovisor-upgradeinfo.json

# Define mounted volume
VOLUME /halo

# Copy binaries from build stages.
COPY --from=build-cosmovisor /ko-app/cosmovisor /usr/local/bin/cosmovisor
COPY --from=build-0-genesis /app /halovisor/genesis/bin/halo
COPY --from=build-1-uluwatu /app /halovisor/upgrades/1_uluwatu/bin/halo

HEALTHCHECK CMD ["/halovisor/upgrades/1_uluwatu/bin/halo", "ready"]

# Cosmovisor is the entrypoint
ENTRYPOINT [ "cosmovisor" ]
# First 'run' is cosmovisor command, second 'run' is halo command.
CMD [ "run", "run" ]
