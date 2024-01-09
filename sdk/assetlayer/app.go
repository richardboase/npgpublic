package assetlayer

type App struct {
	AppID         string  `json:"appId"`
	HandcashAppID string  `json:"handcashAppId"`
	AppName       string  `json:"appName"`
	AppImage      string  `json:"appImage"`
	AppBanner     string  `json:"appBanner"`
	Description   string  `json:"description"`
	URL           string  `json:"url"`
	AutoGrantRead bool    `json:"autoGrantRead"`
	TeamID        string  `json:"teamId"`
	Status        string  `json:"status"`
	CreatedAt     int64   `json:"createdAt"`
	UpdatedAt     int64   `json:"updatedAt"`
	Slots         []*Slot `json:"slots"`
}
type Expression struct {
	ExpressionName string `json:"expressionName"`
	ExpressionID   string `json:"expressionId"`
	ExpressionType struct {
		ExpressionTypeName   string `json:"expressionTypeName"`
		ExpressionAttributes []struct {
			ExpressionAttributeName string `json:"expressionAttributeName"`
			ExpressionAttributeID   string `json:"expressionAttributeId"`
		} `json:"expressionAttributes"`
		ExpressionTypeID string `json:"expressionTypeId"`
	} `json:"expressionType"`
	Description string `json:"description,omitempty"`
}
