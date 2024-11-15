package types

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cosmos/gogoproto/proto"
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
	offsetsByChainVer := make(map[xchain.ChainVersion][]string)
	for _, header := range headers {
		offset := offsetsByChainVer[header.XChainVersion()]
		if len(offset) < logLimit {
			offset = append(offset, strconv.FormatUint(header.AttestOffset, 10))
		} else if len(offset) == logLimit {
			offset = append(offset, "...")
		} else {
			continue
		}
		offsetsByChainVer[header.XChainVersion()] = offset
	}

	attrs := []any{slog.Int("count", len(offsetsByChainVer))}
	for _, header := range headers {
		attrs = append(attrs, slog.String(
			fmt.Sprintf("%d-%d", header.SourceChainId, header.ConfLevel),
			fmt.Sprint(offsetsByChainVer[header.XChainVersion()]),
		))
	}

	return attrs
}

// VotesFromExtension returns the attestations contained in the vote extension, or false if none or an error.
func VotesFromExtension(voteExtension []byte) (*Votes, bool, error) {
	if len(voteExtension) == 0 {
		return nil, false, nil
	}

	resp := new(Votes)
	if err := proto.Unmarshal(voteExtension, resp); err != nil {
		return nil, false, errors.Wrap(err, "decode vote extension")
	}

	return resp, true, nil
}
