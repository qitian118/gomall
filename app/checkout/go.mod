module github.com/cloudwego/biz-demo/gomall/app/checkout

go 1.23

replace (
github.com/apache/thrift => github.com/apache/thrift v0.13.0
github.com/cloudwego/biz-demo/gomall/rpc-gen => ../rpc_gen
)
require github.com/golang/protobuf v1.5.4 // indirect
