package assetlayer

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty   *resty.Client
	headers map[string]string
	host    string
}

func NewClient(host, secret, token string) *Client {
	return &Client{
		resty: resty.New(),
		headers: map[string]string{
			"appsecret": secret,
			"didtoken":  token,
		},
		host: host,
	}
}

func (client *Client) URL(path string) string {
	return fmt.Sprintf("%s/%s", client.host, path)
}

func (client *Client) NewRequest(path string) *resty.Request {
	return client.resty.R()
}
