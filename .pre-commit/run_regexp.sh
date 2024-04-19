#!/usr/bin/env bash

FILES=$@

function check() {
    grep -HnE "$2" $FILES && printf "\n❌ Regexp check failed: %s\n\n" "$1"
}

check 'Log messages must be capitalised' 'log\.(Error|Warn|Info|Debug)\(ctx, "[[:lower:]]' && exit 1
check 'Error messages must not be capitalised' 'errors\.(New|Wrap)\((err, )?"[[:upper:]]' && exit 1
check 'Rather add secrets to baseline with "make secrets-baseline"' 'pragma: allowlist secret' && exit 1
check 'See Go Guidelines for correct error wrapping' '%w' && exit 1

true
