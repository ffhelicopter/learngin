package main

import (
	"fmt"
)

type Collector struct {
	UserAgent      string
	MaxDepth       int
	AllowedDomains []string
}

type Option func(*Collector)

func main() {
	c := NewCollector(UserAgent("Firfox"), MaxDepth(6), AllowedDomains("www.baidu.com"))
	UserAgent("test")(c)   // 正确方式
	UserAgent("Chrome")(c) // 正确方式
}

func NewCollector(options ...Option) *Collector {
	c := &Collector{}
	c.Init()

	for _, f := range options {

		fmt.Println(f, "处理前:", c)
		f(c)
		fmt.Println(f, "处理后:", c)
	}
	return c
}

func UserAgent(ua string) Option {
	return func(c *Collector) {
		c.UserAgent = ua
		fmt.Println("ua:", c)
	}
}

func MaxDepth(depth int) Option {
	return func(c *Collector) {
		c.MaxDepth = depth
	}
}

func AllowedDomains(domains ...string) Option {
	return func(c *Collector) {
		c.AllowedDomains = domains
	}
}

func (c *Collector) Init() {
	c.UserAgent = "Custom"
	c.MaxDepth = 0
}
