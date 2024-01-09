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

func NewClient(secret, token string) *Client {
	return &Client{
		resty: resty.New(),
		headers: map[string]string{
			"appsecret": secret,
			"didtoken":  token,
		},
		host: "https://api-v2.assetlayer.com",
	}
}

func (client *Client) URL(path string) string {
	return fmt.Sprintf("%s/%s", client.host, path)
}

func (client *Client) NewRequest() *resty.Request {
	r := client.resty.R()
	for k, v := range client.headers {
		r = r.SetHeader(k, v)
	}
	return r
}
