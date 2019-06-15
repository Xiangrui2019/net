package web

import (
	"net/http"
	"time"
)

var (
	zeroTime = time.Unix(0, 0)
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Keys     map[string]interface{}
	done     bool
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	if c.Keys != nil {
		value, exists = c.Keys[key]
	}
	return
}

func (c *Context) String(key string) string {
	if v, ok := c.Keys[key]; !ok {
		return ""
	} else {
		return v.(string)
	}
}

func (c *Context) Int(key string) int {
	if v, ok := c.Keys[key]; !ok {
		return 0
	} else {
		return v.(int)
	}
}

func (c *Context) Time(key string) time.Time {
	if v, ok := c.Keys[key]; !ok {
		return zeroTime
	} else {
		return v.(time.Time)
	}
}

func (c *Context) Done() {
	c.done = true
}
