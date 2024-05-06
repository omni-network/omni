package mockresolver

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	fuzz "github.com/google/gofuzz"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	rand "math/rand/v2"
)

// Implements the graphql.Marshaler interface for the Long type.
type Long uint64

func (l Long) ImplementsGraphQLType(name string) bool {
	return name == "Long"
}

func (l *Long) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		value, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return err
		}
		*l = Long(value)
		return nil
	default:
		return fmt.Errorf("cannot unmarshal Long scalar type from %T", input)
	}
}

func (l Long) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d"`, l)), nil
}

// Define the Go struct for the Status enum type.
type Status string

const (
	StatusFailed  Status = "FAILED"
	StatusPending Status = "PENDING"
	StatusSuccess Status = "SUCCESS"
)

// Define the Go struct for the XMsg type.
type XMsg struct {
	ID               graphql.ID
	Block            XBlock
	DestAddress      common.Address
	DestAddressURL   string
	DestChainID      hexutil.Big
	DestGasLimit     hexutil.Big
	DisplayID        string
	Offset           hexutil.Big
	Receipt          *XReceipt
	SourceAddress    common.Address
	SourceAddressURL string
	SourceBlockTime  graphql.Time
	SourceChainID    hexutil.Big
	Status           Status
	TxHash           common.Hash
	TxHashURL        string
}

// Define the Go struct for the XBlock type.
type XBlock struct {
	ID            graphql.ID
	SourceChainID hexutil.Big
	BlockHeight   hexutil.Big
	BlockHash     common.Hash
	Messages      []XMsg
	Timestamp     graphql.Time
}

// Define the Go struct for the XReceipt type.
type XReceipt struct {
	ID             graphql.ID
	GasUsed        hexutil.Big
	Status         Status
	RelayerAddress common.Address
	SourceChainID  hexutil.Big
	DestChainID    hexutil.Big
	Offset         hexutil.Big
	TxHash         common.Hash
	TxHashURL      string
	Timestamp      graphql.Time
}

// Define the Go struct for the Chain type.
type Chain struct {
	ID      graphql.ID
	ChainID hexutil.Big
	Name    string
}

// Define the Go struct for the XMsgConnection type.
type XMsgConnection struct {
	TotalCount Long
	Edges      []XMsgEdge
	PageInfo   PageInfo
}

// Define the Go struct for the XMsgEdge type.
type XMsgEdge struct {
	Cursor graphql.ID
	Node   XMsg
}

// Define the Go struct for the PageInfo type.
type PageInfo struct {
	HasNextPage bool
	HasPrevPage bool
	TotalPages  Long
	CurrentPage Long
}

// Define the Go struct for the Query type.
type QueryResolver struct {
	XBlocks []XBlock
}

func New() *Resolver {
	// Create the root resolver
	resolver := &Resolver{
		QueryResolver: QueryResolver{
			XBlocks: make([]XBlock, 0),
		},
	}

	statuses := []Status{StatusFailed, StatusPending, StatusSuccess}
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 1)
	var relayerAddress common.Address
	fuzzer.Fuzz(&relayerAddress)

	// Populate XBlocks with random data
	for i := 0; i < 31; i++ {
		log.Printf("Generating random XBlock data for block %d of 30\n", i+1)
		var xblock XBlock

		// Fuzz XBlock properties
		xblock.ID = graphql.ID(relay.MarshalID("XBlock", fmt.Sprintf("%d", i+1)))
		fuzzer.Fuzz(&xblock.SourceChainID)
		fuzzer.Fuzz(&xblock.BlockHeight)
		fuzzer.Fuzz(&xblock.BlockHash)
		fuzzer.Fuzz(&xblock.Timestamp)

		numMsgs := rand.IntN(6) // Generate random number of messages between 0 and 5
		for j := 0; j < numMsgs; j++ {
			var xmsg XMsg

			// Fuzz XMsg properties
			xmsg.ID = relay.MarshalID("XMsg", fmt.Sprintf("%d-%d", i+1, j+1))
			fuzzer.Fuzz(&xmsg.Offset)
			fuzzer.Fuzz(&xmsg.SourceAddress)
			fuzzer.Fuzz(&xmsg.DestAddress)
			fuzzer.Fuzz(&xmsg.DestGasLimit)
			fuzzer.Fuzz(&xmsg.SourceChainID)
			fuzzer.Fuzz(&xmsg.DestChainID)
			fuzzer.Fuzz(&xmsg.TxHash)
			fuzzer.Fuzz(&xmsg.SourceBlockTime)
			xmsg.Block = xblock
			xmsg.SourceAddressURL = fmt.Sprintf("https://sepolia.arbiscan.io/address/%s", xmsg.SourceAddress.Hex())
			xmsg.DestAddressURL = fmt.Sprintf("https://sepolia.arbiscan.io/address/%s", xmsg.DestAddress.Hex())
			xmsg.DisplayID = fmt.Sprintf("%s-%s-%s", &xmsg.SourceChainID, &xmsg.DestChainID, &xmsg.Offset)
			xmsg.Status = statuses[rand.IntN(len(statuses))]
			xmsg.TxHashURL = fmt.Sprintf("https://sepolia.arbiscan.io/tx/%s", xmsg.TxHash.Hex())

			var xreceipt XReceipt

			// Fuzz XReceipt properties
			xreceipt.ID = graphql.ID(relay.MarshalID("XReceipt", fmt.Sprintf("%d-%d", i+1, j+1)))
			fuzzer.Fuzz(&xreceipt.GasUsed)
			xreceipt.RelayerAddress = relayerAddress
			fuzzer.Fuzz(&xreceipt.SourceChainID)
			fuzzer.Fuzz(&xreceipt.DestChainID)
			fuzzer.Fuzz(&xreceipt.Offset)
			fuzzer.Fuzz(&xreceipt.TxHash)
			fuzzer.Fuzz(&xreceipt.Timestamp)
			xreceipt.Status = statuses[rand.IntN(len(statuses))]
			xreceipt.TxHashURL = fmt.Sprintf("https://sepolia.arbiscan.io/tx/%s", xreceipt.TxHash.Hex())

			xmsg.Receipt = &xreceipt

			xblock.Messages = append(xblock.Messages, xmsg)
		}

		resolver.XBlocks = append(resolver.XBlocks, xblock)
	}

	return resolver
}

// Define the root resolver.
type Resolver struct {
	QueryResolver
}

// Implement the xblock query resolver.
func (r *QueryResolver) XBlock(ctx context.Context, args struct{ SourceChainID, Height hexutil.Big }) *XBlock {
	for _, xblock := range r.XBlocks {
		if xblock.SourceChainID.String() == args.SourceChainID.String() && xblock.BlockHeight.String() == args.Height.String() {
			return &xblock
		}
	}
	return nil
}

// Implement the xreceipt query resolver.
func (r *QueryResolver) Xreceipt(ctx context.Context, args struct{ SourceChainID, DestChainID, Offset hexutil.Big }) *XReceipt {
	for _, xblock := range r.XBlocks {
		for _, xmsg := range xblock.Messages {
			if xmsg.SourceChainID.String() == args.SourceChainID.String() && xmsg.DestChainID.String() == args.DestChainID.String() && xmsg.Offset.String() == args.Offset.String() {
				return xmsg.Receipt
			}
		}
	}
	return nil
}

// Implement the xmsg query resolver.
func (r *Resolver) Xmsg(ctx context.Context, args struct{ SourceChainID, DestChainID, Offset hexutil.Big }) *XMsg {
	for _, xblock := range r.XBlocks {
		for _, xmsg := range xblock.Messages {
			if xmsg.SourceChainID.String() == args.SourceChainID.String() && xmsg.DestChainID.String() == args.DestChainID.String() && xmsg.Offset.String() == args.Offset.String() {
				return &xmsg
			}
		}
	}
	return nil
}

type XMsgsArgs struct {
	First  *int32
	After  *graphql.ID
	Last   *int32
	Before *graphql.ID
}

// Implement the xmsg query resolver.
func (r *QueryResolver) Xmsgs(ctx context.Context, args XMsgsArgs) (XMsgConnection, error) {
	var messages []XMsg
	for _, xblock := range r.XBlocks {
		messages = append(messages, xblock.Messages...)
	}

	// default length of items to return
	var numItems int32 = 10

	// Apply pagination
	var start, end, pageNum int
	if args.First != nil && args.Last != nil {
		log.Println("Both first and last arguments are provided. Ignoring last argument.")
	}
	if args.Before != nil && args.After != nil {
		return XMsgConnection{}, fmt.Errorf("cannot provide both before and after arguments")
	}

	cur := &cursor{}

	if args.First != nil {
		numItems = *args.First
		if args.After != nil {
			if err := cur.Decode(*args.After); err != nil {
				return XMsgConnection{}, err
			}
			start = int(cur.ID) + 1
		} else {
			start = 0
		}
		pageNum = int(cur.PageNum) + 1
		end = start + int(numItems)
		if end > len(messages) {
			end = len(messages)
		}
	} else if args.Last != nil {
		numItems = *args.Last
		if args.Before != nil {
			if err := cur.Decode(*args.Before); err != nil {
				return XMsgConnection{}, err
			}
			end = int(cur.ID) - 1
		} else {
			end = len(messages)
		}
		pageNum = int(cur.PageNum) - 1
		start = end - int(numItems)
		if start < 0 {
			start = 0
		}
	}

	fmt.Println("Messages length: ", len(messages))
	fmt.Printf("start: %d, end: %d, pageNum: %d\n", start, end, pageNum)

	// Create the edges
	var edges []XMsgEdge
	for i := start; i < end; i++ {
		cur := &cursor{
			ID:      uint(i),
			PageNum: uint(pageNum),
		}
		cursorID, err := cur.Encode()
		if err != nil {
			return XMsgConnection{}, err
		}
		edges = append(edges, XMsgEdge{
			Cursor: cursorID,
			Node:   messages[i],
		})
	}

	// Create the page info
	pageInfo := PageInfo{
		HasNextPage: end < len(messages),
		HasPrevPage: start > 0,
		TotalPages:  Long(uint64(math.Ceil(float64(len(messages)) / float64(numItems)))),
		CurrentPage: Long(uint64(pageNum)),
	}

	return XMsgConnection{
		TotalCount: Long(uint64(len(messages))),
		Edges:      edges,
		PageInfo:   pageInfo,
	}, nil
}

// Implement the supportedChains query resolver.
func (r *Resolver) SupportedChains(ctx context.Context) []*Chain {
	// TODO: Implement logic to fetch supported chains
	return nil
}

type cursor struct {
	ID      uint
	PageNum uint
}

func (c *cursor) Encode() (graphql.ID, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	res := base64.StdEncoding.EncodeToString(b)

	return graphql.ID(res), nil
}

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
