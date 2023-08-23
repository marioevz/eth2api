package builderapi

import (
	"context"

	"github.com/protolambda/eth2api"
	"github.com/protolambda/zrnt/eth2/beacon/common"
)

// Retrieve the expected upcoming withdrawals for a given state.
func ExpectedWithdrawals(ctx context.Context, cli eth2api.Client, stateId eth2api.StateId, dest *common.Withdrawals) (exists bool, err error) {
	return eth2api.SimpleRequest(ctx, cli, eth2api.FmtGET("/eth/v1/builder/states/%s/expected_withdrawals", stateId.StateId()), eth2api.Wrap(dest))
}
