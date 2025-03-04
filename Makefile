.PHONY gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ..\..\idl\ --type RPC --module github.com/gomall_project/demo/demo_proto --service demo_proto --idl ..\..\idl\echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_proto && cwgo server --type RPC --module github.com/gomall_project/demo/demo_thrift --service demo_thrift --idl ..\..\idl\echo.thrift


gen-frontend:
	@cd frontend && cwgo server --module github.com/gomall_project/frontend --service frontend --idl ../idl/echo.proto -I ../..
