package cursors

import "github.com/omni-network/omni/lib/xchain"

func (c *Cursor) StreamID() xchain.StreamID {
	return xchain.StreamID{
		SourceChainID: c.GetSrcChainId(),
		DestChainID:   c.GetDstChainId(),
		ShardID:       xchain.ShardID(c.GetConfLevel()), // todo check
	}
}

func (c *Cursor) Empty() bool {
	return c.GetFirstXmsgOffset() == 0
}
