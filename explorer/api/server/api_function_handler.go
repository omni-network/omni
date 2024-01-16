package server

import (
	"context"
	"log"
	"sync"

	apifunctions "github.com/omni-network/omni/explorer/api/api_functions"
	service "github.com/omni-network/omni/explorer/api/openapi"
)

type RestService struct {
	mux sync.Mutex
}

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

func (s *RestService) HealthCheck(
	ctx context.Context,
) (*service.HealthCheckResponse, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	log.Printf("health check called")

	h, err := apifunctions.GetHealth(ctx)
	if err != nil {
		return nil, err
	}

	return h, nil
}
