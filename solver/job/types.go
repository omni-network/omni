package job

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/core/types"
)

func (j *Job) eventLog() (types.Log, bool, error) {
	if len(j.GetDeprecatedEventJson()) == 0 {
		return types.Log{}, false, nil
	}

	var elog types.Log
	if err := json.Unmarshal(j.GetDeprecatedEventJson(), &elog); err != nil {
		return types.Log{}, false, errors.Wrap(err, "unmarshal event log")
	}

	return elog, true, nil
}

// Migrate updates the job table to use the new event format.
// TODO(corver): Remove this along with the DeprecatedEventJson field.
func Migrate(ctx context.Context, db *DB) (int, error) {
	iter, err := db.table.List(ctx, JobPrimaryKey{})
	if err != nil {
		return 0, errors.Wrap(err, "list jobs")
	}

	var count int
	for iter.Next() {
		job, err := iter.Value()
		if err != nil {
			return 0, errors.Wrap(err, "value")
		}

		elog, ok, err := job.eventLog()
		if err != nil {
			return 0, errors.Wrap(err, "get event log")
		} else if !ok {
			continue // Migration not needed
		} else if len(elog.Topics) == 0 {
			return 0, errors.New("no topics [BUG]")
		}

		orderID, status, err := solvernet.ParseEvent(elog)
		if err != nil {
			return 0, errors.Wrap(err, "parse event")
		}

		statusU64, err := umath.ToUint64(status)
		if err != nil {
			return 0, errors.Wrap(err, "parse status [BUG]")
		}

		job.TxString = elog.TxHash.String()
		job.OrderId = orderID[:]
		job.Status = statusU64
		job.DeprecatedEventJson = nil

		if err := db.table.Update(ctx, job); err != nil {
			return 0, errors.Wrap(err, "save job")
		}

		count++
	}

	return count, nil
}
