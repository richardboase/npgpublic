package assetlayer

import (
	"fmt"

	"github.com/kr/pretty"
)

func (client *Client) Try(method, path string, query map[string]string, body ...interface{}) (interface{}, error) {
	response := &Response{}
	r := client.NewRequest().EnableTrace().SetResult(response)
	if query != nil {
		r = r.SetQueryParams(query)
	}
	if len(body) > 0 {
		r = r.SetBody(body[0])
	}
	url := client.URL(path)
	println(method, url)
	resp, err := r.Execute(method, url)
	if err != nil {
		return nil, err
	}
	code := resp.StatusCode()
	if code < 200 || code >= 300 {
		pretty.Println(response)
		return nil, fmt.Errorf("GetResponse: failed with code: %d", code)
	}
	return response.Body, nil
}
