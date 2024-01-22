package handlers

import (
	"context"

	service "github.com/omni-network/omni/explorer/api/openapi"
)

func GetHealth(_ context.Context) (*service.HealthCheckResponse, error) {
	r := &service.HealthCheckResponse{
		Message: "hello world",
	}

	return r, nil
}
