package grpc

import (
	"context"
	"fmt"

	"github.com/magicsong/kidecar/api"
	"github.com/magicsong/kidecar/pkg/server"
)

// Config 插件配置
type Config struct {
    Port int `json:"port"`
}

type grpc struct {
    config *Config
    mgr    api.SidecarManager
    stopCh chan struct{}
}

func NewPlugin() api.Plugin {
    return &grpc{
        stopCh: make(chan struct{}),
    }
}

func (g *grpc) Name() string {
    return "grpc"
}

func (g *grpc) Version() string {
    return "v1"
}

func (g *grpc) Init(config interface{}, mgr api.SidecarManager) error {
    if c, ok := config.(*Config); ok {
        g.config = c
        g.mgr = mgr
        return nil
    }
    return fmt.Errorf("invalid config type")
}

func (g *grpc) Start(ctx context.Context, errCh chan<- error) {
    // 启动 GRPC 服务器
    go server.StartGRPCServer(g.stopCh)
}

func (g *grpc) Stop(ctx context.Context) error {
    close(g.stopCh)
    return nil
}

func (g *grpc) Status() (*api.PluginStatus, error) {
    return &api.PluginStatus{
        Name:    g.Name(),
        Version: g.Version(),
        Running: true,
    }, nil
}

func (g *grpc) GetConfigType() interface{} {
    return &Config{}
}