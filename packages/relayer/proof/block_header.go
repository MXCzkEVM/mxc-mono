package proof

import (
	"context"

	"github.com/MXCzkEVM/mxc-mono/packages/relayer/encoding"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// blockHeader fetches block via rpc, then converts an ethereum block to the BlockHeader type that LibBridgeData
// uses in our contracts
func (p *Prover) blockHeader(ctx context.Context, blockHash common.Hash) (encoding.BlockHeader, error) {
	h, err := p.blocker.BlockByHash(ctx, blockHash)
	if err != nil {
		return encoding.BlockHeader{}, errors.Wrap(err, "p.ethClient.GetBlockByNumber")
	}

	return encoding.BlockToBlockHeader(h), nil
}
