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
