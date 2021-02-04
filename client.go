package motanx

import (
	"fmt"
	gsync "github.com/snail007/gmc/util/sync"
	"github.com/weibocom/motan-go"
	"github.com/weibocom/motan-go/core"
	"github.com/weibocom/motan-go/protocol"
	"math/rand"
	"sync"
)

var client = NewMClient(5000)

func Call(address string, authToken string, service, method string, args ...interface{}) (reply interface{}, err error) {
	return client.Call(address, authToken, service, method, args...)
}

func CallTimeout(address string, authToken string, service, method string, timeoutMilliseconds int, args ...interface{}) (reply interface{}, err error) {
	return client.CallTimeout(address, authToken, service, method, timeoutMilliseconds, args...)
}

func Close() {
	client.Close()
}

type MClient struct {
	defaultTimeoutMilliseconds int
	epMap                      *sync.Map
}

func NewMClient(defaultTimeoutMilliseconds int) *MClient {
	return &MClient{
		defaultTimeoutMilliseconds: defaultTimeoutMilliseconds,
		epMap:                      &sync.Map{},
	}
}

func (s *MClient) Close() {
	s.epMap.Range(func(key, value interface{}) bool {
		value.(core.EndPoint).Destroy()
		s.epMap.Delete(key)
		return true
	})
}

func (s *MClient) Call(address string, authToken string, service, method string, args ...interface{}) (reply interface{}, err error) {
	return s.CallTimeout(address, authToken, service, method, 0, args...)
}

func (s *MClient) CallTimeout(address string, authToken string, service, method string, timeoutMilliseconds int, args ...interface{}) (reply interface{}, err error) {
	if args == nil {
		args = []interface{}{}
	}
	key := address + service
	var ep core.EndPoint
	if v, exists := s.epMap.Load(key); exists {
		ep = v.(core.EndPoint)
	} else {
		gsync.OnceDo(key, func() {
			clientExt := motan.GetDefaultExtFactory()
			u := core.FromExtInfo(fmt.Sprintf("motan2://%s/%s?serialization=simple&maxRequestTimeout=300000", address, service))
			ep = clientExt.GetEndPoint(u)
			if ep == nil {
				err = fmt.Errorf("build endpoint fail")
				return
			}
			s.epMap.Store(key, ep)
			ep.SetSerialization(core.GetSerialization(u, clientExt))
			core.Initialize(ep)
		})
	}
	if err != nil {
		return
	}
	request := &core.MotanRequest{}
	request.RequestID = rand.Uint64()
	request.ServiceName = service
	request.Method = method
	request.Attachment = core.NewStringMap(core.DefaultAttachmentSize)
	request.Arguments = args
	if timeoutMilliseconds > 0 {
		request.Attachment.Store(protocol.MTimeout, fmt.Sprintf("%d", timeoutMilliseconds))
	}
	if authToken != "" {
		request.Attachment.Store("token", authToken)
	}
	resp := ep.Call(request)
	if resp.GetException() != nil {
		err = fmt.Errorf("call fail, exception: %s, code: %d, tyep: %d",
			resp.GetException().ErrMsg,
			resp.GetException().ErrCode,
			resp.GetException().ErrType,
		)
		return
	}
	reply = resp.GetValue()
	return
}
