package motanx

import (
	"bytes"
	"fmt"
	"github.com/weibocom/motan-go"
	"github.com/weibocom/motan-go/config"
	"github.com/weibocom/motan-go/core"
)

type ServiceItem struct {
	Service interface{}
	Sid     string
}

type MotanServer struct {
	msContext   *motan.MSContext
	mFactory    core.ExtensionFactory
	mConfig     *config.Config
	servicesArr []ServiceItem
	port        string
	authToken   string
}

func (s *MotanServer) Services() []ServiceItem {
	return s.servicesArr
}

func (s *MotanServer) AddService(service interface{}, sid string) {
	s.servicesArr = append(s.servicesArr, ServiceItem{
		Service: service,
		Sid:     sid,
	})
}

func (s *MotanServer) MsContext() *motan.MSContext {
	return s.msContext
}

func (s *MotanServer) MFactory() core.ExtensionFactory {
	return s.mFactory
}

func (s *MotanServer) MConfig() *config.Config {
	return s.mConfig
}

func NewMotanServer(port, authToken string) (s *MotanServer) {
	s = &MotanServer{
		port:      port,
		authToken: authToken,
	}
	return
}

func (s *MotanServer) Start() (err error) {
	var conf *config.Config
	cfg := `
motan-server:
  log_dir: "stdout"
  application: "app"
motan-registry:
  direct:
    protocol: direct
motan-service:
`
	serviceTpl := `  %s:
    path: %s
    protocol: motan2
    registry: direct
    serialization: simple
    ref: %s
    filter: "auth"
    requestTimeout: 600000
    export: "motan2:%s"
`
	s.AddService(&helloService{}, "motanx.hello")
	for _, v := range s.servicesArr {
		cfg += fmt.Sprintf(serviceTpl, v.Sid, v.Sid, v.Sid, s.port)
	}
	conf, err = config.NewConfigFromReader(bytes.NewReader([]byte(cfg)))
	if err != nil {
		return
	}
	s.mFactory = motan.GetDefaultExtFactory()
	s.msContext = motan.NewMotanServerContextFromConfig(conf)
	s.mConfig = conf
	s.mFactory.RegistExtFilter("auth", func() core.Filter {
		return GetAuthFilter(s.authToken)
	})
	for _, v := range s.servicesArr {
		s.msContext.RegisterService(v.Service, v.Sid)
	}
	s.msContext.Start(s.mFactory)
	s.msContext.ServicesAvailable()
	reply, err := CallTimeout("127.0.0.1:"+s.port, s.authToken, "motanx.hello", "Ping", 1000)
	if err != nil {
		return
	}
	if !reply.(bool) {
		return fmt.Errorf("start montan server fail, response false")
	}
	return
}

type helloService struct{}

func (s *helloService) Ping() bool {
	return true
}
