package assetlayer

type Collection struct {
	CollectionName   string   `json:"collectionName"`
	Description      string   `json:"description"`
	Type             string   `json:"type"`
	SlotID           string   `json:"slotId"`
	Maximum          int      `json:"maximum"`
	Tags             []string `json:"tags"`
	RoyaltyRecipient string   `json:"royaltyRecipient"`
	Properties       struct {
		Prop1 int `json:"prop1"`
		Prop2 struct {
			Val1 int `json:"val1"`
			Val2 int `json:"val2"`
		} `json:"prop2"`
	} `json:"properties"`
	CollectionImage string `json:"collectionImage"`
}

func (client *Client) NewCollection() (string, error) {
	collection := &Collection{}
	return collection.CollectionName, nil
}
