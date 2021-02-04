package tests

import (
	"github.com/snail007/motanx"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	assert := assert2.New(t)
	server := motanx.NewMotanServer("33880", "")
	server.AddService(&HelloService{}, "hello")
	err := server.Start()
	assert.Nil(err)
}

func TestCall(t *testing.T) {
	assert := assert2.New(t)
	server := motanx.NewMotanServer("33880", "token_secret")
	server.AddService(&HelloService{}, "hello")
	err := server.Start()
	assert.Nil(err)
	reply, err := motanx.Call("127.0.0.1:33880","token_secret", "hello", "Hello", "jack")
	assert.Nil(err)
	assert.Equal("hello jack", reply)
}

type HelloService struct {
}

func (s *HelloService) Hello(name string) string {
	return "hello " + name
}
