package assetlayer

type Response struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Body       interface{} `json:"body"`
}
