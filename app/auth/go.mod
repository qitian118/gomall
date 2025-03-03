module github.com/cloudwego/biz-demo/gomall/app/auth

go 1.24.0

replace (
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
	github.com/cloudwego/biz-demo/gomall/rpc_gen => ../../rpc_gen
)
require (
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
)
