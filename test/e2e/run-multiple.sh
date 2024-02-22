#!/usr/bin/env bash
#
# This is a convenience script that takes a list of testnet manifests
# as arguments and runs each one of them sequentially. If a testnet
# fails, the container logs are dumped to stdout along with the testnet
# manifest, but the remaining testnets are still run.
#
# This is mostly used to run generated networks in nightly CI jobs.
#

set -euo pipefail

if [[ $# == 0 ]]; then
	echo "Usage: $0 [MANIFEST...]" >&2
	exit 1
fi

echo "🌊==> Running e2e tests:" "$@"

for MANIFEST in "$@"; do
	START=$SECONDS
	echo "🌊==> Running manifest: $MANIFEST"

	if ! e2e -f "$MANIFEST"; then
		echo "🌊==> ❌ Testnet $MANIFEST failed, dumping manifest..."
		cat "$MANIFEST"

		echo "🌊==> Dumping failed container logs to failed-logs.txt..."
		e2e -f "$MANIFEST" logs > failed-logs.txt

		echo "🌊==> Displaying failed container error and warn logs..."
		grep -iE "(panic|erro|warn)" failed-logs.txt || echo "No errors or warns found"

		echo "🌊==> Cleaning up failed manifest $MANIFEST..."
		e2e -f "$MANIFEST" clean

    echo "🌊==> ❌ Manifest $MANIFEST failed..."
		exit 1
	fi

	echo "🌊==> ✅ Completed manifest $MANIFEST in $(( SECONDS - START ))s"
	echo ""
done

echo "🌊==> 🎉 All manifests successful "
