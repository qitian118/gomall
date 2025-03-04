package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/gomall_project/app/frontend/conf"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CheckoutClient checkoutservice.Client
	once           sync.Once
	CartClient     cartservice.Client
)

func Init() {
	once.Do(func() {
		iniUserClient()
		initProductClient()

	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}
