package ext

import (
	"github.com/omni-network/omni/e2e/ext/flags"
	"github.com/omni-network/omni/e2e/ext/gaspumps"
	"github.com/omni-network/omni/e2e/ext/omnibridge"
	"github.com/omni-network/omni/e2e/ext/solvernet"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
)

func FromRegex(r flags.Regex) ([]types.Extension, error) {
	flags, err := r.Resolve()
	if err != nil {
		return nil, err
	}

	var exts []types.Extension
	for _, f := range flags {
		ext, err := FromFlag(f)
		if err != nil {
			return nil, err
		}
		exts = append(exts, ext)
	}

	return exts, nil
}

func FromFlag(f flags.Flag) (types.Extension, error) {
	switch f {
	case flags.SolverNet:
		return solvernet.Ext(), nil
	case flags.GasPumps:
		return gaspumps.Ext(), nil
	case flags.OMNIBridge:
		return omnibridge.Ext(), nil
	default:
		return nil, errors.New("unknown flag", "flag", f)
	}
}
