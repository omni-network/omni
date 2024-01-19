package server

import (
	"context"

	apifunctions "github.com/omni-network/omni/explorer/api/api_functions"
	service "github.com/omni-network/omni/explorer/api/openapi"
	"github.com/omni-network/omni/lib/log"
)

type RestService struct{}

func CreateRestService(
	_ context.Context,
) *RestService {
	return &RestService{}
}

// NewError implements api.Handler.
func (*RestService) NewError(
	_ context.Context,
	err error,
) *service.ErrorStatusCode {
	return &service.ErrorStatusCode{
		StatusCode: 501,
		Response: service.Error{
			Code:    500,
			Message: err.Error(),
		},
	}
}

func (*RestService) HealthCheck(
	ctx context.Context,
) (*service.HealthCheckResponse, error) {
	log.Debug(ctx, "Health check called")

	h, err := apifunctions.GetHealth(ctx)
	if err != nil {
		return nil, err
	}

	return h, nil
}
