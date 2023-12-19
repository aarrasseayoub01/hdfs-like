module github.com/aarrasseayoub01/hdfs-mini

replace github.com/aarrasseayoub01/namenode/protobuf/hdfs => ../protobuf

go 1.21.1

require (
	github.com/aarrasseayoub01/namenode/protobuf/hdfs v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/stretchr/testify v1.8.4
	google.golang.org/grpc v1.60.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
