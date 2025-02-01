package flags

import (
	"regexp"

	"github.com/omni-network/omni/lib/errors"
)

type Regex string
type Flag string

const (
	// SolverNet identifies the solver-net e2e extension.
	SolverNet Flag = "solvernet"
	// GasPumps identifies the gas-pumps e2e extension.
	GasPumps Flag = "gaspumps"
	// OMNIBridge identifies the omni-bridge e2e extension.
	OMNIBridge Flag = "omnibridge"
)

func All() []Flag {
	return []Flag{SolverNet, GasPumps, OMNIBridge}
}

func (r Regex) Resolve() ([]Flag, error) {
	// if empty, match all
	p := string(r)
	if p == "" {
		p = ".*"
	}

	pattern, err := regexp.Compile(p)
	if err != nil {
		return nil, errors.Wrap(err, "compile regex")
	}

	var flags []Flag
	for _, flag := range All() {
		if pattern.MatchString(string(flag)) {
			flags = append(flags, flag)
		}
	}

	if len(flags) == 0 {
		return nil, errors.New("no flags matched by regex", "regex", r)
	}

	return flags, nil
}
