package assetlayer

import "fmt"

func (client *Client) Try(method, path string, body ...interface{}) (interface{}, error) {
	response := &Response{}
	r := client.NewRequest().SetResult(response)
	if len(body) > 0 {
		r = r.SetBody(body[0])
	}
	resp, err := r.Execute(method, client.URL(path))
	if err != nil {
		return nil, err
	}
	code := resp.StatusCode()
	if code < 200 || code >= 300 {
		return nil, fmt.Errorf("GetResponse: failed with code: %d", code)
	}
	return response.Body, nil
}
