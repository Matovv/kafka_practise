package containers

import (
	"context"
	"local/service1/internal/pkg/configs"
	"local/service1/internal/pkg/http"
)

type Container struct {
	server *http.Server
}

func NewContainer(configPath string) (*Container, error) {
	config, err := configs.ParseConfig(configPath)
	if err != nil {
		return nil, err
	}
	router := http.NewRouter(config)
	server := http.NewServer(config, router.Handler())
	return &Container{
		server: server,
	}, nil
}

func (c *Container) Start(ctx context.Context) error {
	if err := c.server.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (c *Container) Stop(ctx context.Context) error {
	if err := c.server.Stop(ctx); err != nil {
		return err
	}
	return nil
}