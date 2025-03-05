package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CheckoutClient checkoutservice.Client
	AuthClient     authservice.Client
	OrderClient    orderservice.Client
	ServiceName    = frontendUtils.ServiceName
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	once           sync.Once
	CartClient     cartservice.Client
	err            error
)

func Init() {
	once.Do(func() {
		iniUserClient()
		initProductClient()
		initCheckoutClient()
		initAuthClient()
		initOrderClient()
		initCartClient()
	})
}

func iniUserClient() {
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initAuthClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	AuthClient, err = authservice.NewClient("auth", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	// frontendUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}
