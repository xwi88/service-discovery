module github.com/xwi88.com/service-discovery

go 1.14

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.3
	github.com/google/uuid v1.1.1 // indirect
	github.com/satori/go.uuid v1.2.0
	go.etcd.io/etcd v3.3.20+incompatible
	go.uber.org/zap v1.15.0 // indirect
	google.golang.org/genproto v0.0.0-20200424135956-bca184e23272 // indirect
	google.golang.org/grpc v1.27.0

)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
