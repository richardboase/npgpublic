package assetlayer

import "encoding/json"

type Collection struct {
	CollectionName   string                 `json:"collectionName"`
	Description      string                 `json:"description"`
	Type             string                 `json:"type"`
	SlotID           string                 `json:"slotId"`
	Maximum          int                    `json:"maximum"`
	Tags             []string               `json:"tags"`
	RoyaltyRecipient string                 `json:"royaltyRecipient"`
	Properties       map[string]interface{} `json:"properties"`
	CollectionImage  string                 `json:"collectionImage"`
}

func (client *Client) NewCollection() (string, error) {
	collection := &Collection{}
	return collection.CollectionName, nil
}

func (client *Client) GetCollections(slotID string) ([]*Collection, error) {

	data, err := client.Try(
		"GET",
		"/api/v1/slot/collections",
		map[string]string{
			"slotId": slotID,
			"idOnly": "false",
		},
	)
	if err != nil {
		return nil, err
	}

	m, err := assertMapStringInterface(data)
	if err != nil {
		return nil, err
	}
	app, err := assertMapStringInterface(m["slot"])
	if err != nil {
		return nil, err
	}
	s, err := assertInterfaceArray(app["collections"])
	if err != nil {
		return nil, err
	}

	collections := []*Collection{}

	for _, item := range s {
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		collection := &Collection{}
		if err := json.Unmarshal(b, collection); err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}

	return collections, nil
}
