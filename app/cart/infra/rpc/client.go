package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/qitian118/gomall/app/cart/conf"
	cartutils "github.com/qitian118/gomall/app/cart/utils"
	"github.com/qitian118/gomall/common/clientsuite"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	// r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// cartutils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
