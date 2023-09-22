package beaconapi

import (
	"context"

	"github.com/protolambda/eth2api"
	"github.com/protolambda/zrnt/eth2/beacon/deneb"
)

// Retrieves block details for given block id.
func BlobSidecars(ctx context.Context, cli eth2api.Client, blockId eth2api.BlockId, dest *[]deneb.BlobSidecar) (exists bool, err error) {
	return eth2api.SimpleRequest(ctx, cli, eth2api.FmtGET("/eth/v1/beacon/blob_sidecars/%s", blockId.BlockId()), eth2api.Wrap(dest))
}
