export ROOT_MOD=github.com/cloudwego/biz-demo/gomall
.PHONY:gen_frontend
gen_frontend:cwgo server --type HTTP --idl ../../idl/frontend/cart.proto --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --module ${ROOT_MOD}/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --module  ${ROOT_MOD}/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/order_page.proto

.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC  --service user --module  ${ROOT_MOD}/app/user  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/user.proto
# 添加了go mod init 使其定位到包名
.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/order.proto
	@cd app/order && go mod init github.com/cloudwego/biz-demo/gomall/app/order && cwgo server --type RPC  --service order --module  ${ROOT_MOD}/app/order  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-payment
gen-user:
	@cd app/payment && cwgo server --type RPC  --service payment --module  ${ROOT_MOD}/app/payment  --pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/payment.proto


.PHONY: gen-cart
gen-cart: 
	@cd rpc_gen && cwgo client --type RPC --service cart --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module ${ROOT_MOD}/app/cart --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto

.PHONY: gen-order
gen-order: 
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module ${ROOT_MOD}/app/order --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-auth
gen-auth: 
	@cd rpc_gen && cwgo client --type RPC --service auth --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/auth.proto
	@cd app/auth && cwgo server --type RPC --service auth --module ${ROOT_MOD}/app/auth --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/auth.proto
