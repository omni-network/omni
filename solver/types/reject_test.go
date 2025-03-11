package types_test

import (
	"testing"

	"github.com/omni-network/omni/solver/types"

	"github.com/stretchr/testify/require"
)

func TestRejectReasonValues(t *testing.T) {
	t.Parallel()

	// Ensure that reject reason int values do not change, since they are persisted on-chain.
	require.EqualValues(t, 1, types.RejectDestCallReverts)
	require.EqualValues(t, 2, types.RejectInvalidDeposit)
	// ...
	require.EqualValues(t, 11, types.RejectExpenseOverMax)
	require.EqualValues(t, 12, types.RejectExpenseUnderMin)
}
