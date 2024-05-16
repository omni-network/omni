package data

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/graphql/uintconv"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (p Provider) XMsgCount(ctx context.Context) (*hexutil.Big, bool, error) {
	query, err := p.cl.Msg.Query().Count(ctx)
	if err != nil {
		return nil, false, err
	}

	hex, err := uintconv.ToHex(uint64(query))
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding block count")
	}

	return &hex, true, nil
}

func (p Provider) XMsg(ctx context.Context, srcChainID, destChainID, offset uint64) (*XMsg, bool, error) {
	query, err := p.cl.Msg.Query().
		Where(
			msg.SourceChainID(srcChainID),
			msg.DestChainID(destChainID),
			msg.Offset(offset),
		).
		First(ctx)
	if err != nil {
		return nil, false, err
	}

	block := query.QueryBlock().OnlyX(ctx)
	receipts := query.QueryReceipts().AllX(ctx)

	res, err := EntMsgToGraphQLXMsgWithEdges(ctx, query)
	if err != nil {
		return nil, false, errors.Wrap(err, "decoding message")
	}

	var receipt *XReceipt
	var rts time.Time
	for _, r := range receipts {
		rec, err := EntReceiptToGraphQLXReceipt(ctx, r, block)
		if err != nil {
			return nil, false, errors.Wrap(err, "decoding receipt")
		}
		// If there are multiple receipts use the latest. Timestamp check for the first receipt
		// found would result to true since all receipts are created after time.Time{}.
		if r.CreatedAt.After(rts) {
			rts = r.CreatedAt
			receipt = rec
		}
	}

	res.Receipt = receipt

	res.Chainer = p.ch
	res.Block.Chainer = p.ch
	if res.Receipt != nil {
		res.Receipt.Chainer = p.ch
	}

	return res, true, nil
}

type dbXMsgRow struct {
	ID                    uint64
	Sender                []byte
	To                    []byte
	Data                  []byte
	GasLimit              int
	SourceChainID         uint64
	DestChainID           uint64
	Offset                uint64
	TxHash                common.Hash
	Status                string
	CreatedAt             time.Time
	ReceiptID             sql.NullInt64
	ReceiptTxHash         *common.Hash
	ReceiptSuccess        sql.NullBool
	ReceiptRelayerAddress *[]byte
	ReceiptGasUsed        sql.NullInt64
	ReceiptSourceChainID  sql.NullInt64
	ReceiptDestChainID    sql.NullInt64
	ReceiptOffset         sql.NullInt64
	ReceiptCreatedAt      sql.NullTime
	BlockID               uint64
	BlockChainID          uint64
	BlockHash             []byte
	BlockHeight           uint64
	BlockTimestamp        time.Time
}

func (r *dbXMsgRow) ToGraphQLXMsg(ch Chainer) (*XMsg, error) {
	b := &XBlock{
		Chainer: ch,

		ID:        relay.MarshalID("xblock", r.BlockID),
		ChainID:   hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.BlockChainID))),
		Hash:      common.Hash(r.BlockHash),
		Height:    hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.BlockHeight))),
		Timestamp: graphql.Time{Time: r.BlockTimestamp},
	}

	m := XMsg{
		Chainer: ch,

		ID:            relay.MarshalID("xmsg", r.ID),
		Data:          hexutil.Bytes(r.Data),
		GasLimit:      hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.GasLimit))),
		SourceChainID: hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.SourceChainID))),
		DestChainID:   hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.DestChainID))),
		Offset:        hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.Offset))),
		TxHash:        r.TxHash,
		Status:        Status(r.Status),
		Block:         *b,
	}
	m.DisplayID = fmt.Sprintf("%d-%d-%d", m.SourceChainID.ToInt().Int64(), m.DestChainID.ToInt().Int64(), m.Offset.ToInt().Int64())

	if r.ReceiptSuccess.Valid {
		m.Receipt = &XReceipt{
			Chainer: ch,

			ID:      relay.MarshalID("xreceipt", r.ReceiptID),
			TxHash:  *r.ReceiptTxHash,
			Success: r.ReceiptSuccess.Bool,
			Relayer: common.Address(*r.ReceiptRelayerAddress),
		}
	}
	if r.ReceiptGasUsed.Valid {
		m.Receipt.GasUsed = hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.ReceiptGasUsed.Int64)))
	}
	if r.ReceiptSourceChainID.Valid {
		m.Receipt.SourceChainID = hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.ReceiptSourceChainID.Int64)))
	}
	if r.ReceiptDestChainID.Valid {
		m.Receipt.DestChainID = hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.ReceiptDestChainID.Int64)))
	}
	if r.ReceiptOffset.Valid {
		m.Receipt.Offset = hexutil.Big(*hexutil.MustDecodeBig(fmt.Sprintf("0x%x", r.ReceiptOffset.Int64)))
	}

	return &m, nil
}

type XMsgFilters struct {
	SourceChainID *uint64
	DestChainID   *uint64
	Addr          *common.Address
	TxHash        *common.Hash
	Status        *Status
}

func (p Provider) XMsgs(ctx context.Context, first, last *int32, before *graphql.ID, after *graphql.ID, filters *XMsgFilters) (XMsgConnection, error) {
	if first != nil && last != nil {
		return XMsgConnection{}, errors.New("first and last are mutually exclusive")
	}
	if before != nil && after != nil {
		return XMsgConnection{}, errors.New("before and after are mutually exclusive")
	}

	total := p.StatsProvider.TotalMsgs() // cached value of the total messages
	var (
		numItems   int32
		cur        cursor
		pageNum    int
		totalPages int
		start      uint64
		end        uint64
	)
	if first != nil {
		numItems = *first
	} else if last != nil {
		numItems = *last
	} else {
		return XMsgConnection{}, errors.New("either first or last must be provided")
	}

	totalPages = int(math.Ceil(float64(total) / float64(numItems)))

	// our data is backwards (oldest data is first), so we need to reverse the order
	if first != nil {
		if after != nil {
			if err := cur.Decode(*after); err != nil {
				return XMsgConnection{}, err
			}
			start = cur.ID + 1
			pageNum = int(cur.PageNum) - 1
		} else {
			start = 0
			pageNum = totalPages
		}
		end = start + uint64(numItems)
		if end > total {
			end = total
		}
	} else if last != nil {
		if before != nil {
			if err := cur.Decode(*before); err != nil {
				return XMsgConnection{}, err
			}
			end = cur.ID // exclusive
			pageNum = int(cur.PageNum) + 1
		} else {
			end = total
			pageNum = 1
		}

		if end < uint64(numItems) { // end - numItems < 0
			start = 0
		} else {
			start = end - uint64(numItems)
		}
	}

	log.Info(ctx, "XMsgs", "start", start, "end", end, "total", total, "numItems", numItems, "pageNum", pageNum, "totalPages", totalPages)

	query := `
	SELECT
		m.id,
		m.sender,
		m.to,
		m.data,
		m.gas_limit,
		m.source_chain_id,
		m.dest_chain_id,
		m.offset,
		m.tx_hash,
		m.status,
		m.created_at,
		r.id AS receipt_id,
		r.tx_hash AS receipt_tx_hash,
		r.success AS receipt_success,
		r.relayer_address AS receipt_relayer_address,
		r.gas_used AS receipt_gas_used,
		r.source_chain_id AS receipt_source_chain_id,
		r.dest_chain_id AS receipt_dest_chain_id,
		r.offset AS receipt_offset,
		r.created_at AS receipt_created_at,
		b.id AS block_id,
		b.chain_id AS block_chain_id,
		b.hash AS block_hash,
		b.height AS block_height,
		b.timestamp AS block_timestamp
	FROM
		msgs m
			LEFT JOIN block_msgs bm ON bm.msg_id = m.id
			LEFT JOIN blocks b ON bm.block_id = b.id
			LEFT JOIN msg_receipts mr ON mr.msg_id = m.id
			LEFT JOIN (
				-- use only the last receipt row per tx_hash
				SELECT
					*
				FROM
					receipts
				WHERE id IN
					(
						SELECT
							MAX(r1.id) AS id
						FROM
							receipts r1
						GROUP BY
							r1.tx_hash
					)
			) r ON mr.receipt_id = r.id
	WHERE
		m.id <= $2
		AND m.status = 'SUCCESS'
		--AND m.source_chain_id = 1651
		AND m.dest_chain_id = 1654
		-- AND m.tx_hash = decode('22f3133cdc633e1eaa163e28c383f94267d5c26a5cfb7b86946ba696a3bc0ab6', 'hex')
	ORDER BY
		m.id DESC -- the id column is auto-increment and therefore chronological
	LIMIT $1;
	`
	rows, err := p.cl.DB().QueryContext(ctx, query, numItems, end)
	if err != nil {
		return XMsgConnection{}, errors.Wrap(err, "preparing query")
	}
	defer rows.Close()

	var res XMsgConnection
	res.TotalCount = Long(total)

	for rows.Next() {
		var v dbXMsgRow

		err := rows.Scan(
			&v.ID,
			&v.Sender,
			&v.To,
			&v.Data,
			&v.GasLimit,
			&v.SourceChainID,
			&v.DestChainID,
			&v.Offset,
			&v.TxHash,
			&v.Status,
			&v.CreatedAt,
			&v.ReceiptID,
			&v.ReceiptTxHash,
			&v.ReceiptSuccess,
			&v.ReceiptRelayerAddress,
			&v.ReceiptGasUsed,
			&v.ReceiptSourceChainID,
			&v.ReceiptDestChainID,
			&v.ReceiptOffset,
			&v.ReceiptCreatedAt,
			&v.BlockID,
			&v.BlockChainID,
			&v.BlockHash,
			&v.BlockHeight,
			&v.BlockTimestamp,
		)
		if err != nil {
			return XMsgConnection{}, errors.Wrap(err, "scanning row")
		}

		m, err := v.ToGraphQLXMsg(p.ch)
		if err != nil {
			return XMsgConnection{}, errors.Wrap(err, "decoding message")
		}
		edge := XMsgEdge{
			Cursor: relay.MarshalID("cursor", m.ID),
			Node:   *m,
		}
		res.Edges = append(res.Edges, edge)
	}

	res.PageInfo = PageInfo{
		HasNextPage: pageNum < totalPages,
		HasPrevPage: pageNum > 1,
		TotalPages:  Long(uint64(totalPages)),
		CurrentPage: Long(uint64(pageNum)),
	}

	return res, nil
}

// cursor is used for data pagination and is supposed to be represented as a base64 encoded json object which would allow for
// future refactoring of the pagination without breaking the API or require frontend changes.
type cursor struct {
	ID      uint64 `json:"id"`
	PageNum uint64 `json:"page_num"`
}

// Encode encodes the cursor to a base64 string.
func (c *cursor) Encode() (graphql.ID, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	res := base64.StdEncoding.EncodeToString(b)

	return graphql.ID(res), nil
}

// Decode decodes the cursor from a base64 string.
func (c *cursor) Decode(id graphql.ID) error {
	if len(id) == 0 {
		return nil
	}
	b, err := base64.StdEncoding.DecodeString(string(id))
	if err != nil {
		return err
	}

	return json.Unmarshal(b, c)
}
