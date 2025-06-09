.PHONY: build
build: # Build localnet/devnet solver programs
	# Ensure the localnet keypair is used
	mkdir -p ./target/deploy
	cp ./localnet/solver_inbox-keypair.json ./target/deploy/solver_inbox-keypair.json
	# Build
	anchor build
	anchor keys sync
	# Copy the compiled artefacts to localnet
	cp ./target/idl/solver_inbox.json ./localnet/solver_inbox.json
	cp ./target/deploy/solver_inbox.so ./localnet/solver_inbox.so
	cp ./target/types/solver_inbox.ts ./localnet/solver_inbox.ts
	# Generate go bindings
	solana-anchor-go -src=./target/idl/solver_inbox.json -pkg=anchorinbox -dst=./anchorinbox -remove-account-suffix
	# Fix go bindings (replace .ParamsOrderId with .Params.OrderId)
	sed -i.bak 's/\.ParamsOrderId/\.Params.OrderId/g' anchorinbox/open.go
	rm anchorinbox/open.go.bak
