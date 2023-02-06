module github.com/snail007/motanx

go 1.13

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/snail007/gmc v0.0.0-20230206035746-0feadd2a7a90
	github.com/snail007/go-sqlcipher v0.0.0-20210114093415-fb27975e042f // indirect
	github.com/stretchr/testify v1.7.0
	github.com/weibocom/motan-go v0.0.0-20210129094724-c208e12f05e7
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace (
	go.uber.org/atomic => github.com/uber-go/atomic v1.4.0
	go.uber.org/multierr => github.com/uber-go/multierr v1.1.1-0.20180122172545-ddea229ff1df
	go.uber.org/zap => github.com/uber-go/zap v1.9.1
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/net => github.com/golang/net v0.0.0-20181017193950-04a2e542c03f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181011152604-fa43e7bc11ba
	golang.org/x/text => github.com/golang/text v0.3.0
	google.golang.org/appengine => github.com/golang/appengine v1.2.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20181016170114-94acd270e44e
	google.golang.org/grpc => github.com/grpc/grpc-go v1.15.0
)
