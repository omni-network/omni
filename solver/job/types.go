package job

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/core/types"
)

func (j *Job) EventLog() (types.Log, error) {
	var elog types.Log
	if err := json.Unmarshal(j.GetEventJson(), &elog); err != nil {
		return types.Log{}, errors.Wrap(err, "unmarshal event log")
	}

	return elog, nil
}
