package client

import (
	"fmt"
	"prometheusd/config"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

// http通信
type Client struct {
	option  *config.Option
	request *req.Req
}

// 构造对象函数
func NewClient(option *config.Option) *Client {
	return &Client{
		option:  option,
		request: req.New(),
	}
}

// 注册
func (c *Client) Register(event map[string]interface{}) {
	path := fmt.Sprintf("%s/agent/register/", c.option.Server)
	respoonse, err := c.request.Post(path, req.BodyJSON(event))
	// fmt.Println(respoonse)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(respoonse)
}

// 心跳
func (c *Client) Heartbeat(event map[string]interface{}) {
	path := fmt.Sprintf("%s/agent/heartbeat/", c.option.Server)
	response, err := c.request.Get(path, req.Param(event))
	// fmt.Println(response)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Debug(response)
}

// config
func (c *Client) Config(event map[string]interface{}) (string, error) {
	path := fmt.Sprintf("%s/agent/config/", c.option.Server)
	respoonse, err := c.request.Get(path, req.Param(event))
	// fmt.Println(respoonse)
	if err != nil {
		return "", nil
	}
	return respoonse.ToString()
}
