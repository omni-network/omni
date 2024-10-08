package indexer

import (
	"bytes"
	"encoding/json"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func (b *Block) XChainBlock() (xchain.Block, error) {
	var resp xchain.Block
	if err := json.Unmarshal(b.GetBlockJson(), &resp); err != nil { //nolint:musttag // TODO: Rather use protobuf
		return xchain.Block{}, errors.Wrap(err, "unmarshal block")
	}

	return resp, nil
}

func (l *MsgLink) Hash() (common.Hash, error) {
	return cast.EthHash(l.GetIdHash())
}

// IsMsg returns true if the message belongs to this link.
func (l *MsgLink) IsMsg(msg xchain.Msg) bool {
	return bytes.Equal(l.GetIdHash(), msg.Hash().Bytes())
}

// IsReceipt returns true if the receipt belongs to this link.
func (l *MsgLink) IsReceipt(receipt xchain.Receipt) bool {
	return bytes.Equal(l.GetIdHash(), receipt.Hash().Bytes())
}
