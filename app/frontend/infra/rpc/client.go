package rpc

import (
	"context"
	"log"
	"sync"

	"github.com/YiD11/gomall/common/clientsuite"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/user/userservice"
	consulclient "github.com/kitex-contrib/config-consul/client"
	"github.com/kitex-contrib/config-consul/consul"

	"github.com/YiD11/gomall/app/frontend/conf"
	frontendUtils "github.com/YiD11/gomall/app/frontend/utils"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client

	currentServiceName = conf.GetConf().Hertz.Service
	registryAddr       = conf.GetConf().Registry.RegistryAddress

	once sync.Once
	err  error
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	var opts []client.Option

	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: currentServiceName, RegistryAddr: registryAddr}))

	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleErr(err)
}

func initProductClient() {
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)

	fbp := fallback.NewFallbackPolicy(
		fallback.UnwrapHelper(
			func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
				if err == nil {
					return resp, nil
				}
				methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
				if methodName != "ListProducts" {
					return resp, err
				}
				return &product.ListProductsResp{
					Products: []*product.Product{
						{
							Price:       6.6,
							Id:          3,
							Picture:     "/static/image/t-shirt-1.jpeg",
							Name:        "T-shirt",
							Description: "This is a T-shirt",
						},
					},
				}, nil
			}),
	)

	consulClient, err := consul.NewClient(consul.Options{
		Addr: conf.GetConf().Registry.RegistryAddress[0],
	})
	if err != nil {
		log.Panicln(err)
	}

	var opts []client.Option

	opts = append(opts,
		client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: currentServiceName, RegistryAddr: registryAddr}),
		client.WithCircuitBreaker(cbs),
		client.WithFallback(fbp),
		client.WithSuite(consulclient.NewSuite("product", currentServiceName, consulClient)),
	)

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleErr(err)
}

func initCartClient() {
	var opts []client.Option

	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: currentServiceName, RegistryAddr: registryAddr}))

	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleErr(err)
}

func initCheckoutClient() {
	var opts []client.Option

	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: currentServiceName, RegistryAddr: registryAddr}))

	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleErr(err)
}

func initOrderClient() {
	var opts []client.Option

	opts = append(opts, client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: currentServiceName, RegistryAddr: registryAddr}))

	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleErr(err)
}
