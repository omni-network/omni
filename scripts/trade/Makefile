.PHONY: build-docker
build-docker:
	@GORELEASER_CURRENT_TAG=v0.0.2 \
	GORELEASER_PREVIOUS_TAG=v0.0.1 \
	goreleaser release -f .goreleaser-snapshot.yaml --snapshot --clean --skip=archive
