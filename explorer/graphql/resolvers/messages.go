package resolvers

import (
	"context"

	"github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

const MsgsLimit = 25

type MessagesProvider interface {
	XMsg(ctx context.Context, sourceChainID uint64, destChainID uint64, offset uint64) (*data.XMsg, bool, error)
	XMsgs(ctx context.Context, first, last *int32, before *graphql.ID, after *graphql.ID, filters *data.XMsgFilters) (data.XMsgConnection, error)
}

type MessagesResolver struct {
	p MessagesProvider
}

func NewMessagesResolver(p MessagesProvider) *MessagesResolver {
	return &MessagesResolver{
		p: p,
	}
}

type XMsgRangeArgs struct {
	From hexutil.Big
	To   hexutil.Big
}

type XMsgArgs struct {
	SourceChainID hexutil.Big
	DestChainID   hexutil.Big
	Offset        hexutil.Big
}

func (m *MessagesResolver) XMsg(ctx context.Context, args XMsgArgs) (*data.XMsg, error) {
	res, found, err := m.p.XMsg(ctx, args.SourceChainID.ToInt().Uint64(), args.DestChainID.ToInt().Uint64(), args.Offset.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch message")
	}
	if !found {
		return nil, errors.New("message not found")
	}

	return res, nil
}

type XMsgsArgs struct {
	Filters *[]FilterInput
	First   *int32
	After   *graphql.ID
	Last    *int32
	Before  *graphql.ID
}

type FilterInput struct {
	Key   string
	Value string
}

func (a *XMsgsArgs) DataFilters() (data.XMsgFilters, error) {
	var res data.XMsgFilters
	if a.Filters != nil {
		for _, f := range *a.Filters {
			switch f.Key {
			case "srcChainID":
				val := hexutil.MustDecodeBig(f.Value)
				res.SourceChainID = uint64Ptr(val.Uint64())
			case "destChainID":
				val := hexutil.MustDecodeBig(f.Value)
				res.DestChainID = uint64Ptr(val.Uint64())
			case "address":
				val := common.HexToAddress(f.Value)
				res.Addr = &val
			case "txHash":
				val := common.HexToHash(f.Value)
				res.TxHash = &val
			case "status":
				val, err := data.ParseStatus(f.Value)
				if err != nil {
					return res, errors.New("invalid status")
				}
				res.Status = &val
			}
		}
	}

	return res, nil
}

func (a *XMsgsArgs) Validate() error {
	if a.First == nil && a.Last == nil {
		return errors.New("either first or last must be provided")
	}
	if a.First != nil && a.Last != nil {
		return errors.New("first and last are mutually exclusive")
	}
	if a.Before != nil && a.After != nil {
		return errors.New("before and after are mutually exclusive")
	}
	if a.First != nil {
		if *a.First < 0 {
			return errors.New("first must be positive")
		}
		if *a.First > MsgsLimit {
			return errors.New("first exceeds limit")
		}
	}
	if a.Last != nil {
		if *a.Last < 0 {
			return errors.New("last must be positive")
		}
		if *a.Last > MsgsLimit {
			return errors.New("last exceeds limit")
		}
	}

	return nil
}

func (m *MessagesResolver) XMsgs(ctx context.Context, args XMsgsArgs) (data.XMsgConnection, error) {
	if err := args.Validate(); err != nil {
		return data.XMsgConnection{}, err
	}
	filters, err := args.DataFilters()
	if err != nil {
		return data.XMsgConnection{}, err
	}

	return m.p.XMsgs(ctx, args.First, args.Last, args.Before, args.After, &filters)
}

func uint64Ptr(v uint64) *uint64 {
	return &v
}
