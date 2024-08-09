# Docker build args
# TODO: Pin halo genesis to public omega release
ARG HALO_VERSION_GENESIS=main

# Build stages
FROM golang:alpine AS build-cosmovisor
ENV COSMOVISOR_VERSION=v1.5.0
RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@${COSMOVISOR_VERSION}

FROM omniops/halo:${HALO_VERSION_GENESIS} AS build-genesis

# TODO: Add post-genesis network-upgrade versions here.

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
ENV COSMOVISOR_TIMEFORMAT_LOGS=rfc3339nano

# Define mounted volumes
VOLUME /halo/config
VOLUME /halo/data

# Copy binaries from build stages.
COPY --from=build-cosmovisor /go/bin/cosmovisor /usr/local/bin/cosmovisor
COPY --from=build-genesis /app $DAEMON_HOME/cosmovisor/genesis/bin/halo

# Cosmovisor is the entrypoint
ENTRYPOINT [ "cosmovisor" ]
# First 'run' is cosmovisor command, second 'run' is halo command.
CMD [ "run", "run" ]
