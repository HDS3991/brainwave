package client

import (
	"brainwave/pkg/grpc/service"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"sync/atomic"
	"time"
)

// grpc客户端连接
// 没做同步，需要提前初始化
var clients map[service.Name]*GrpcClient
var initialed bool

func InitClients(cfgs ...ClientConfig) {

	if initialed {
		panic("只能初始化一次！")
	}
	initialed = true

	clients = make(map[service.Name]*GrpcClient, len(cfgs))
	for _, cfg := range cfgs {
		cli := NewGrpcClient(cfg)
		clients[cfg.ServiceName] = cli
	}
}

// GetConn
//
// @@Description：获取服务对应的客户端连接
//
// @@param serviceName：服务名
func GetConn(ctx context.Context, serviceName service.Name) (conn *grpc.ClientConn, err error) {

	cli, ok := clients[serviceName]
	if !ok {

		return nil, errors.New(fmt.Sprintf("client not found, service: %s", serviceName))
	}

	return cli.GetConn(ctx)
}

// 客户端配置
type ClientConfig struct {
	ServiceName service.Name // 服务名
	ServiceAddr string       // 服务地址
}

type GrpcClient struct {
	cfg  ClientConfig
	conn *grpc.ClientConn
	done uint32
	m    sync.Mutex
}

func NewGrpcClient(cfg ClientConfig) *GrpcClient {
	return &GrpcClient{cfg: cfg}
}

func (cli *GrpcClient) GetConn(ctx context.Context) (conn *grpc.ClientConn, err error) {

	if atomic.LoadUint32(&cli.done) == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		err = cli.initSlow(ctx)
	}

	return cli.conn, err
}

func (cli *GrpcClient) initSlow(ctx context.Context) (err error) {

	cli.m.Lock()
	defer cli.m.Unlock()
	if cli.done == 0 {
		err = cli.initConn(ctx)
		if err == nil {
			atomic.StoreUint32(&cli.done, 1)
		}
	}

	return
}

func (cli *GrpcClient) initConn(ctx context.Context) (err error) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	cli.conn, err = grpc.DialContext(ctx,
		cli.cfg.ServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), //建立真实连接
		grpc.WithChainUnaryInterceptor(),
	)
	if err != nil {
		err = errors.New(fmt.Sprintf("did not connect: %v", err))
	}

	return
}
