package server

import (
	"context"
	"net/http"

	service "github.com/omni-network/omni/explorer/api/openapi"
)

type RestServerClient interface {
	CreateServer(ctx context.Context, port int) *http.Server
}

type ClientImpl struct {
	RestServerClient
	port  int
	title string
}

func NewClient(port int) *ClientImpl {
	client := ClientImpl{
		port:  port,
		title: "rest server",
	}

	return &client
}

func (ClientImpl) CreateServer(ctx context.Context) (http.Handler, error) {
	apis := CreateRestService(ctx)

	s, err := service.NewServer(apis)
	if err != nil {
		return nil, err
	}

	return s, nil
}
