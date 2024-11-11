package types

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/lib/xchain"
)

const logLimit = 5

// VoteLogs returns the logs as opinionated human-readable logging attributes.
func VoteLogs(votes []*Vote) []any {
	var headers []*AttestHeader
	for _, vote := range votes {
		headers = append(headers, vote.GetAttestHeader())
	}

	return AttLogs(headers)
}

// AttLogs returns the headers as opinionated human-readable logging attributes.
func AttLogs(headers []*AttestHeader) []any {
	offsets := make(map[xchain.ChainVersion][]string)
	for _, header := range headers {
		offset := offsets[header.XChainVersion()]
		if len(offset) < logLimit {
			offset = append(offset, strconv.FormatUint(header.AttestOffset, 10))
		} else if len(offset) == logLimit {
			offset = append(offset, "...")
		} else {
			continue
		}
		offsets[header.XChainVersion()] = offset
	}

	attrs := []any{slog.Int("count", len(offsets))}
	for chainVer, offsets := range offsets {
		attrs = append(attrs, slog.String(
			fmt.Sprintf("%d-%d", chainVer.ID, chainVer.ConfLevel),
			fmt.Sprint(offsets),
		))
	}

	return attrs
}
