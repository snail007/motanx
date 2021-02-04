module github.com/snail007/motanx

go 1.13

require (
	github.com/snail007/gmc v0.0.0-20210201024958-51dc8b862fe3
	github.com/stretchr/testify v1.7.0
	github.com/weibocom/motan-go v0.0.0-20210129094724-c208e12f05e7
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
