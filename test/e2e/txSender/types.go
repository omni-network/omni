package txsender

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/xchain"
)

func MsgToBindings(xcall xchain.Msg) bindings.XTypesMsg {
	return bindings.XTypesMsg{
		SourceChainId: xcall.SourceChainID,
		DestChainId:   xcall.DestChainID,
		StreamOffset:  xcall.StreamOffset,
		Sender:        xcall.SourceMsgSender,
		To:            xcall.DestAddress,
		Data:          xcall.Data,
		GasLimit:      xcall.DestGasLimit,
	}
}
