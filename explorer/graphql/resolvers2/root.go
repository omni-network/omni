package resolvers2

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	fuzz "github.com/google/gofuzz"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

//go:embed index.html
var indexPage []byte

//go:embed schema.graphql
var schema string

// Define the Go struct for the Status enum type
type Status string

const (
	StatusFailed  Status = "FAILED"
	StatusPending Status = "PENDING"
	StatusSuccess Status = "SUCCESS"
)

// Define the Go struct for the XMsg type
type XMsg struct {
	ID                  graphql.ID
	DisplayID           string
	Offset              hexutil.Big
	SourceMessageSender common.Address
	DestAddress         common.Address
	DestGasLimit        hexutil.Big
	SourceChainID       hexutil.Big
	DestChainID         hexutil.Big
	TxHash              common.Hash
	BlockHeight         hexutil.Big
	BlockHash           common.Hash
	ReceiptTxHash       common.Hash
	Status              Status
	SourceBlockTime     graphql.Time
	Block               XBlock
	Receipt             *XReceipt
}

// Define the Go struct for the XBlock type
type XBlock struct {
	ID            graphql.ID
	SourceChainID hexutil.Big
	BlockHeight   hexutil.Big
	BlockHash     common.Hash
	Messages      []XMsg
	Timestamp     graphql.Time
}

// Define the Go struct for the XReceipt type
type XReceipt struct {
	ID             graphql.ID
	GasUsed        hexutil.Big
	Status         Status
	RelayerAddress common.Address
	SourceChainID  hexutil.Big
	DestChainID    hexutil.Big
	Offset         hexutil.Big
	TxHash         common.Hash
	Timestamp      graphql.Time
	Block          XBlock
	Message        XMsg
}

// Define the Go struct for the Chain type
type Chain struct {
	ID      graphql.ID
	ChainID hexutil.Big
	Name    string
}

// Define the Go struct for the XMsgConnection type
type XMsgConnection struct {
	TotalCount hexutil.Big
	Edges      []*XMsgEdge
	PageInfo   PageInfo
}

// Define the Go struct for the XMsgEdge type
type XMsgEdge struct {
	Cursor graphql.ID
	Node   *XMsg
}

// Define the Go struct for the PageInfo type
type PageInfo struct {
	HasNextPage bool
	HasPrevPage bool
	TotalPages  int
	CurrentPage int
}

// Define the Go struct for the Query type
type Query struct {
	XBlocks          []*XBlock
	XBlockResolver   *XBlockResolver
	XReceiptResolver *XReceiptResolver
	XMsgResolver     *XMsgResolver
}

// Define the Go struct for the XBlockResolver type
type XBlockResolver struct {
	XBlocks []*XBlock
}

// Define the Go struct for the XReceiptResolver type
type XReceiptResolver struct {
	XReceipts []*XReceipt
}

// Define the Go struct for the XMsgResolver type
type XMsgResolver struct {
	XMsgs []*XMsg
}

// Define the root resolver
type Resolver struct {
	Query
}

// Implement the xblock query resolver
func (r *Resolver) XBlock(ctx context.Context, args struct{ SourceChainID, Height hexutil.Big }) *XBlock {
	for _, xblock := range r.XBlocks {
		if xblock.SourceChainID == args.SourceChainID && xblock.BlockHeight == args.Height {
			return xblock
		}
	}
	return nil
}

// Implement the xreceipt query resolver
func (r *Resolver) Xreceipt(ctx context.Context, args struct{ SourceChainID, DestChainID, Offset hexutil.Big }) *XReceipt {
	for _, xblock := range r.XBlocks {
		for _, xmsg := range xblock.Messages {
			if xmsg.SourceChainID == args.SourceChainID && xmsg.DestChainID == args.DestChainID && xmsg.Offset == args.Offset {
				return xmsg.Receipt
			}
		}
	}
	return nil
}

// Implement the xmsg query resolver
func (r *Resolver) Xmsg(ctx context.Context, args struct{ SourceChainID, DestChainID, Offset hexutil.Big }) *XMsg {
	for _, xblock := range r.XBlocks {
		for _, xmsg := range xblock.Messages {
			if xmsg.SourceChainID == args.SourceChainID && xmsg.DestChainID == args.DestChainID && xmsg.Offset == args.Offset {
				return xmsg
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

// Implement the xmsgs query resolver
func (r *Resolver) Xmsgs(ctx context.Context, args XMsgsArgs) (XMsgConnection, error) {
	var messages []XMsg
	for _, xblock := range r.XBlocks {
		messages = append(messages, xblock.Messages...)
	}

	// Apply pagination
	var start, end int
	var err error
	if args.First != nil {
		start = int(*args.First)
	} else if args.Last != nil {
		start = len(messages) - int(*args.Last)
		if start < 0 {
			start = 0
		}
	}
	if args.After != nil {
		cursor := *args.After
		start, err = strconv.Atoi(string(cursor)) + 1
	}
	if args.Before != nil {
		cursor := relay.FromGlobalID(string(*args.Before))
		if cursor != nil {
			end = cursor.(int)
		}
	}
	if args.First != nil && args.Last != nil {
		log.Println("Both first and last arguments are provided. Ignoring last argument.")
	}
	if args.First != nil {
		end = start + int(*args.First)
		if end > len(messages) {
			end = len(messages)
		}
	} else if args.Last != nil {
		start = end - int(*args.Last)
		if start < 0 {
			start = 0
		}
	}

	// Create the edges
	var edges []*XMsgEdge
	for i := start; i < end; i++ {
		edges = append(edges, &XMsgEdge{
			Cursor: graphql.ID(relay.ToGlobalID("XMsg", fmt.Sprintf("%d", i))),
			Node:   messages[i],
		})
	}

	// Create the page info
	pageInfo := PageInfo{
		HasNextPage: end < len(messages),
		HasPrevPage: start > 0,
		TotalPages:  len(messages),
		CurrentPage: start + 1,
	}

	return &XMsgConnection{
		TotalCount: hexutil.Big{Value: fmt.Sprintf("%d", len(messages))},
		Edges:      edges,
		PageInfo:   pageInfo,
	}
}

// Implement the supportedChains query resolver
func (r *Resolver) SupportedChains(ctx context.Context) []*Chain {
	// TODO: Implement logic to fetch supported chains
	return nil
}

func main() {
	// Create the root resolver
	resolver := &Resolver{
		Query: Query{
			XBlocks: make([]*XBlock, 0),
			XBlockResolver: &XBlockResolver{
				XBlocks: make([]*XBlock, 0),
			},
			XReceiptResolver: &XReceiptResolver{
				XReceipts: make([]*XReceipt, 0),
			},
			XMsgResolver: &XMsgResolver{
				XMsgs: make([]*XMsg, 0),
			},
		},
	}

	// Populate XBlocks with random data
	fuzzer := fuzz.New()
	for i := 0; i < 50; i++ {
		xblock := &XBlock{
			ID:            graphql.ID(relay.ToGlobalID("XBlock", fmt.Sprintf("%d", i))),
			SourceChainID: hexutil.Big{Value: fmt.Sprintf("%d", i)},
			BlockHeight:   hexutil.Big{Value: fmt.Sprintf("%d", i)},
			BlockHash:     common.Hash{Value: fmt.Sprintf("0x%x", i)},
			Messages:      make([]*XMsg, 0),
			Timestamp:     Time{Value: "2022-01-01T00:00:00Z"},
		}

		// Populate XMsgs with random data
		for j := 0; j < 4; j++ {

			var xmsg []XMsg

			var xreceipt XReceipt
			fuzzer.Fuzz(&xreceipt)

			xmsg.Receipt = &xreceipt

			xblock.Messages = append(xblock.Messages, xmsg)
			resolver.XMsgResolver.XMsgs = append(resolver.XMsgResolver.XMsgs, xmsg)
			resolver.XReceiptResolver.XReceipts = append(resolver.XReceiptResolver.XReceipts, xreceipt)
		}

		resolver.XBlocks = append(resolver.XBlocks, xblock)
		resolver.XBlockResolver.XBlocks = append(resolver.XBlockResolver.XBlocks, xblock)
	}

	// Create the GraphQL schema
	gqlSchema := graphql.MustParseSchema(schema, &Resolver{})

	http.HandleFunc("/", home)

	// Serve the GraphQL API
	http.Handle("/graphql", &relay.Handler{Schema: gqlSchema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write(indexPage)
}
