package types

import (
	"context"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
)

type Result struct {
	OrderID solvernet.OrderID
	Data    bindings.SolverNetOrderData
}

type Job struct {
	// Name is the friendly name of the job
	Name string

	// Cadence is interval at which to run the job
	Cadence time.Duration

	// SrcChainID is chain ID of the inbox.
	SrcChainID uint64

	// OpenOrdersFunc opens multiple orders and returns their results.
	// Note it may open no orders.
	OpenOrdersFunc func(ctx context.Context) ([]Result, error)
}
