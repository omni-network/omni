package data

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/graphql/uintconv"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type RowQuerier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

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
	RowNum                uint64
	ID                    uint64
	Sender                common.Address
	To                    common.Address
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
	ReceiptRelayerAddress *common.Address
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
		Sender:        r.Sender,
		To:            r.To,
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
			Relayer: *r.ReceiptRelayerAddress,
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
		numItems int32
		cur      cursor
		beforeID uint64
		afterID  uint64
	)
	if first != nil {
		numItems = *first
	} else if last != nil {
		numItems = *last
	} else {
		return XMsgConnection{}, errors.New("either first or last must be provided")
	}

	// our data is backwards (oldest data is first), so we need to reverse the order
	if first != nil {
		if after != nil {
			if err := cur.Decode(*after); err != nil {
				return XMsgConnection{}, err
			}
			afterID = cur.ID
		}
	} else if last != nil {
		if before != nil {
			if err := cur.Decode(*before); err != nil {
				return XMsgConnection{}, err
			}
			beforeID = cur.ID
		}
	}

	var status string
	if filters.Status != nil {
		status = string(*filters.Status)
	}
	var srcChainID, destChainID uint64
	if filters.SourceChainID != nil {
		srcChainID = *filters.SourceChainID
	}
	if filters.DestChainID != nil {
		destChainID = *filters.DestChainID
	}

	query := `
	SELECT
		*
	FROM
		(
			SELECT DISTINCT ON (m.id)
				ROW_NUMBER() OVER (ORDER BY m.id DESC) AS row_num, -- increasing row number for pagination purposes of data in reverse order (highest id first)
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
				r.relayer_address::bytea AS receipt_relayer_address,
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
					LEFT JOIN receipts r ON mr.receipt_id = r.id
			WHERE
				($4 = '' OR m.status = $4)                                  -- status filter
				AND ($5 = 0 OR m.source_chain_id = $5)                      -- source_chain_id filter
				AND ($6 = 0 OR m.dest_chain_id = $6)                        -- dest_chain_id filter
				AND ($7::bytea IS NULL OR m.tx_hash = $7 OR r.tx_hash = $7) -- tx_hash filter
				AND ($8::bytea IS NULL OR m.sender = $8 OR m.to = $8)       -- address filter
			ORDER BY
				m.id DESC, -- the id column is auto-increment and desc returns latest data first
				r.id DESC  -- the id column is auto-increment - together with distinct ensures that only the latest receipt is returned
		) AS x
	WHERE
		($2 = 0 OR x.row_num > $2)     -- before cursor
	    AND ($3 = 0 OR x.row_num < $3) -- after cursor
	LIMIT $1
	OFFSET $9
	`
	var queryOffset uint64
	if first != nil {
		// since the data is ordered by id in descending order, we need to adjust the offset for forwards pagination (e.g. going to the previous page)
		if afterID > uint64(numItems) {
			queryOffset = afterID - uint64(numItems)
		} else {
			queryOffset = 0
		}
	}
	rows, err := p.cl.DB().QueryContext(ctx, query, numItems, beforeID, afterID, status, srcChainID, destChainID, filters.TxHash, filters.Addr, queryOffset)
	if err != nil {
		return XMsgConnection{}, errors.Wrap(err, "preparing query")
	}
	defer rows.Close()

	var res XMsgConnection
	res.TotalCount = Long(total)

	var firstRowNum uint64
	var firstPopulated bool
	for rows.Next() {
		if rows.Err() != nil {
			return XMsgConnection{}, errors.Wrap(rows.Err(), "executing query")
		}
		var v dbXMsgRow

		err := rows.Scan(
			&v.RowNum,
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
		if !firstPopulated {
			firstRowNum = v.RowNum
			firstPopulated = true
		}

		m, err := v.ToGraphQLXMsg(p.ch)
		if err != nil {
			return XMsgConnection{}, errors.Wrap(err, "decoding message")
		}
		cur := cursor{ID: v.RowNum}
		cv, err := cur.Encode()
		if err != nil {
			return XMsgConnection{}, errors.Wrap(err, "encoding cursor")
		}
		edge := XMsgEdge{
			Cursor: cv,
			Node:   *m,
		}
		res.Edges = append(res.Edges, edge)
	}
	totalCount, err := msgsCount(ctx, p.cl.DB(), srcChainID, destChainID, status, filters.Addr, filters.TxHash)
	if err != nil {
		return XMsgConnection{}, errors.Wrap(err, "query total messages count")
	}

	res.TotalCount = Long(totalCount)
	totalPages := uint64(math.Ceil(float64(totalCount) / float64(numItems)))
	pageNum := uint64(math.Ceil(float64(firstRowNum) / float64(numItems)))
	res.PageInfo = PageInfo{
		HasNextPage: pageNum < totalPages,
		HasPrevPage: pageNum > 1,
		TotalPages:  Long(totalPages),
		CurrentPage: Long(pageNum),
	}

	return res, nil
}

func msgsCount(ctx context.Context, q RowQuerier, srcChainID, destChainID uint64, status string, addr *common.Address, txHash *common.Hash) (uint64, error) {
	query := `
	SELECT
		COUNT(DISTINCT m.id)
	FROM
		msgs m
			LEFT JOIN msg_receipts mr ON mr.msg_id = m.id
			LEFT JOIN receipts r ON mr.receipt_id = r.id
	WHERE
		($1 = '' OR m.status = $1)                                  -- status filter
		AND ($2 = 0 OR m.source_chain_id = $2)                      -- source_chain_id filter
		AND ($3 = 0 OR m.dest_chain_id = $3)                        -- dest_chain_id filter
		AND ($4::bytea IS NULL OR m.tx_hash = $4 OR r.tx_hash = $4) -- tx_hash filter
		AND ($5::bytea IS NULL OR m.sender = $5 OR m.to = $5)       -- address filter
	`

	var count uint64
	err := q.QueryRowContext(ctx, query, status, srcChainID, destChainID, txHash, addr).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "query row")
	}

	return count, nil
}

// cursor is used for data pagination and is supposed to be represented as a base64 encoded json object which would allow for
// future refactoring of the pagination without breaking the API or require frontend changes.
type cursor struct {
	ID uint64 `json:"id"`
}

// Encode encodes the cursor to a base64 string.
func (c *cursor) Encode() (graphql.ID, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", errors.Wrap(err, "encoding cursor")
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
		return errors.Wrap(err, "decoding cursor")
	}
	err = json.Unmarshal(b, c)
	if err != nil {
		return errors.Wrap(err, "invalid cursor")
	}

	return nil
}
