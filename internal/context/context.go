package context

import "net/http"

type ProxyContext interface {
	Handle()
	Reset()
}

// Context 用于Proxy的上下文传递

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

func (c *Context) Handle() {
	panic("not implemented") // TODO: Implement
}

func (c *Context) Reset() {
	c.Request = nil
	c.Request = nil
}
