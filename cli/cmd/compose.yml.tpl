services:

  halo:
    container_name: halo
    image: omniops/halovisor:{{.HaloTag}}
    restart: unless-stopped
    ports:
      - 26656:26656 # Consensus P2P
      - 26657:26657 # Consensus RPC
      # - 26660:26660 # Prometheus metrics
    volumes:
      - ./halo/config:/halo/config
      - ./halo/data:/halo/data
      - ./geth/geth/jwtsecret:/geth/jwtsecret

  omni_evm:
    container_name: omni_evm
    image: ethereum/client-go:{{.GethTag}}
    restart: unless-stopped
    command:
      - --config=/geth/config.toml
      - --verbosity={{.GethVerbosity}}                 # Log level (1=error,2=warn,3=info,4=debug)
      {{ if .GethArchive }}- --gcmode=archive{{ end }}
      # Flags not available via config.toml
      #- --nat=extip:<my-external-ip> # External IP for P2P via NAT
      #- --metrics                    # Enable prometheus metrics
      #- --pprof                      # Enable prometheus metrics
      #- --pprof.addr=0.0.0.0         # Enable prometheus metrics

    ports:
      - 8551         # Auth-RPC (used by halo)
      - 8545:8545    # JSON-RCP
      - 8546:8546    # Websocket-RPC
      - 30303:30303  # Execution P2P
      #- 6060:6060   # Prometheus metrics
    volumes:
      - ./geth:/geth
