module gvisor.dev/gvisor

go 1.16

require (
	cloud.google.com/go/bigquery v1.8.0
	github.com/BurntSushi/toml v0.3.1
	github.com/bazelbuild/rules_go v0.27.0
	github.com/cenkalti/backoff v1.1.1-0.20190506075156-2146c9339422
	github.com/containerd/cgroups v1.0.3
	github.com/containerd/console v1.0.3
	github.com/containerd/containerd v1.6.1
	github.com/containerd/fifo v1.0.0
	github.com/containerd/go-runc v1.0.0
	github.com/containerd/typeurl v1.0.2
	github.com/docker/docker v1.4.2-0.20191028175130-9e7d5ac5ea55
	github.com/docker/go-connections v0.4.0
	github.com/gofrs/flock v0.8.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/google/btree v1.0.1
	github.com/google/go-cmp v0.5.6
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20210423192551-a2663126120b
	github.com/google/subcommands v1.0.2-0.20190508160503-636abe8753b8
	github.com/google/uuid v1.2.0
	github.com/kr/pty v1.1.5
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/mohae/deepcopy v0.0.0-20170308212314-bb9b5e7adda9
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/sirupsen/logrus v1.8.1
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635
	github.com/vishvananda/netlink v1.1.1-0.20210330154013-f5de75959ad5
	github.com/xeipuuv/gojsonschema v1.2.0
	go.uber.org/multierr v1.6.0
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	golang.org/x/tools v0.1.5
	google.golang.org/api v0.43.0
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
	honnef.co/go/tools v0.1.1
	k8s.io/api v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
)
