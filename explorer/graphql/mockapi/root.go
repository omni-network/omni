package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"math"
	"math/big"
	rand "math/rand/v2"
	"net/http"

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
	ReceiptTxHash       *common.Hash
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
	Edges      []XMsgEdge
	PageInfo   PageInfo
}

// Define the Go struct for the XMsgEdge type
type XMsgEdge struct {
	Cursor graphql.ID
	Node   XMsg
}

// Define the Go struct for the PageInfo type
type PageInfo struct {
	HasNextPage bool
	HasPrevPage bool
	TotalPages  int32
	CurrentPage int32
}

// Define the Go struct for the Query type
type QueryResolver struct {
	XBlockResolver
	XReceiptResolver
	XMsgResolver

	XBlocks []XBlock
}

// Define the Go struct for the XBlockResolver type
type XBlockResolver struct {
	XBlocks []XBlock
}

// Define the Go struct for the XReceiptResolver type
type XReceiptResolver struct {
	XReceipts []XReceipt
}

// Define the Go struct for the XMsgResolver type
type XMsgResolver struct {
	XMsgs []XMsg
}

// Define the root resolver
type Resolver struct {
	QueryResolver
}

// Implement the xblock query resolver
func (r *QueryResolver) XBlock(ctx context.Context, args struct{ SourceChainID, Height hexutil.Big }) *XBlock {
	for _, xblock := range r.XBlocks {
		if xblock.SourceChainID.String() == args.SourceChainID.String() && xblock.BlockHeight.String() == args.Height.String() {
			return &xblock
		}
	}
	return nil
}

// Implement the xreceipt query resolver
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

// Implement the xmsg query resolver
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

// Implement the xmsg query resolver
func (r *QueryResolver) Xmsgs(ctx context.Context, args XMsgsArgs) (XMsgConnection, error) {
	var messages []XMsg
	for _, xblock := range r.XBlocks {
		messages = append(messages, xblock.Messages...)
	}

	// default length of items to return
	var numItems int32 = 10

	// Apply pagination
	var start, end int
	if args.First != nil && args.Last != nil {
		log.Println("Both first and last arguments are provided. Ignoring last argument.")
	}
	if args.Before != nil && args.After != nil {
		return XMsgConnection{}, fmt.Errorf("cannot provide both before and after arguments")
	}
	if args.First != nil {
		start = 0
		numItems = *args.First
	} else if args.Last != nil {
		start = len(messages) - int(*args.Last)
		if start < 0 {
			start = 0
		}
	} else {
		return XMsgConnection{}, fmt.Errorf("either first or last argument must be provided")
	}
	if args.After != nil {
		var cursor int
		err := relay.UnmarshalSpec(*args.After, &cursor)
		if err != nil {
			return XMsgConnection{}, err
		}
		start = cursor + 1
	}
	if args.Before != nil {
		var cursor int
		err := relay.UnmarshalSpec(*args.After, &cursor)
		if err != nil {
			return XMsgConnection{}, err
		}
		end = cursor
	}
	if args.Before == nil && args.After == nil {
		end = start + int(numItems)
	}

	// Create the edges
	var edges []XMsgEdge
	for i := start; i < end; i++ {
		edges = append(edges, XMsgEdge{
			Cursor: relay.MarshalID("XMsg", i),
			Node:   messages[i],
		})
	}

	// Create the page info
	pageInfo := PageInfo{
		HasNextPage: end < len(messages),
		HasPrevPage: start > 0,
		TotalPages:  int32(math.Ceil(float64(len(messages)) / float64(numItems))),
		CurrentPage: int32(math.Ceil(float64(start)/float64(numItems))) + 1,
	}

	return XMsgConnection{
		TotalCount: hexutil.Big(*big.NewInt(int64(len(messages)))),
		Edges:      edges,
		PageInfo:   pageInfo,
	}, nil
}

// Implement the supportedChains query resolver
func (r *Resolver) SupportedChains(ctx context.Context) []*Chain {
	// TODO: Implement logic to fetch supported chains
	return nil
}

func main() {
	// Create the root resolver
	resolver := &Resolver{
		QueryResolver: QueryResolver{
			XBlocks: make([]XBlock, 0),
			XBlockResolver: XBlockResolver{
				XBlocks: make([]XBlock, 0),
			},
			XReceiptResolver: XReceiptResolver{
				XReceipts: make([]XReceipt, 0),
			},
			XMsgResolver: XMsgResolver{
				XMsgs: make([]XMsg, 0),
			},
		},
	}

	statuses := []Status{StatusFailed, StatusPending, StatusSuccess}
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 1)
	var relayerAddress common.Address
	fuzzer.Fuzz(&relayerAddress)

	// Populate XBlocks with random data
	for i := 0; i < 30; i++ {
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
			fuzzer.Fuzz(&xmsg.SourceMessageSender)
			fuzzer.Fuzz(&xmsg.DestAddress)
			fuzzer.Fuzz(&xmsg.DestGasLimit)
			fuzzer.Fuzz(&xmsg.SourceChainID)
			fuzzer.Fuzz(&xmsg.DestChainID)
			fuzzer.Fuzz(&xmsg.TxHash)
			fuzzer.Fuzz(&xmsg.BlockHeight)
			fuzzer.Fuzz(&xmsg.BlockHash)
			fuzzer.Fuzz(&xmsg.ReceiptTxHash)
			fuzzer.Fuzz(&xmsg.SourceBlockTime)
			xmsg.Block = xblock
			xmsg.DisplayID = fmt.Sprintf("%s-%s-%s", &xmsg.SourceChainID, &xmsg.DestChainID, &xmsg.Offset)
			xmsg.Status = statuses[rand.IntN(len(statuses))]

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

			xmsg.Receipt = &xreceipt
			xblock.Messages = append(xblock.Messages, xmsg)
		}

		resolver.XBlocks = append(resolver.XBlocks, xblock)
		resolver.XBlockResolver.XBlocks = append(resolver.XBlockResolver.XBlocks, xblock)
	}

	opsts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}

	// Create the GraphQL schema
	gqlSchema := graphql.MustParseSchema(schema, &Resolver{}, opsts...)

	log.Println("Server started at http://localhost:8888")
	http.HandleFunc("/", home)
	http.Handle("/query", &relay.Handler{Schema: gqlSchema})
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write(indexPage)
}
