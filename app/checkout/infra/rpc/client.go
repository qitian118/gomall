package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/cloudwego/kitex/client"
	"github.com/qitian118/gomall/app/checkout/conf"
	"github.com/qitian118/gomall/common/clientsuite"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}
func initCartClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// if err != nil {
	// 	panic(err)
	// }
	// opts = append(opts, client.WithResolver(r))
	// opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	// 	client.WithTransportProtocol(transport.GRPC),
	// 	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	// )
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}

}
func initPaymentClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// if err != nil {
	// 	panic(err)
	// }
	// opts = append(opts, client.WithResolver(r))
	// opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
	// 	client.WithTransportProtocol(transport.GRPC),
	// 	client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	// )
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}

}
func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}

}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}

}
