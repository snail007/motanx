# motanx

Motan RPC server & client in easy way.


## Server

```go
package main

import (
	"github.com/snail007/motanx"
)

func main() {
	server := motanx.NewMotanServer("33880", "")
	server.AddService(&HelloService{}, "MyHelloService")
	err := server.Start()
	if err!=nil{
		panic(err)
	}
    select{}
}

type HelloService struct {
}

func (s *HelloService) Hello(name string) string {
	return "hello " + name
}

func (s *HelloService) Ping() bool {
	return true
}

```

# Client

```go
package main

import (
	"fmt"
	"github.com/snail007/motanx"
)

func main() {
	reply, err := motanx.Call("127.0.0.1:33880","", "MyHelloService", "Hello", "jack")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("result: %s",reply.(string))
    reply, err = motanx.Call("127.0.0.1:33880","", "MyHelloService", "Ping")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("result: %s",reply.(bool))

}
```

