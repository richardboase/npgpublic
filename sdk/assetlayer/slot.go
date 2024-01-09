package assetlayer

import (
	"encoding/json"
	"fmt"
)

type Slot struct {
	SlotName             string        `json:"slotName"`
	Description          string        `json:"description"`
	SlotImage            string        `json:"slotImage"`
	AppID                string        `json:"appId"`
	Collections          []string      `json:"collections"`
	Expressions          []*Expression `json:"expressions"`
	AcceptingCollections bool          `json:"acceptingCollections"`
	IsPublic             bool          `json:"isPublic"`
	CollectionTypes      string        `json:"collectionTypes"`
	CreatedAt            int64         `json:"createdAt"`
	UpdatedAt            int64         `json:"updatedAt"`
	SlotID               string        `json:"slotId"`
}

func (client *Client) GetSlots() ([]*Slot, error) {

	data, err := client.Try("GET", "/api/slots")
	if err != nil {
		return nil, err
	}
	app, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to assert type: %v", data)
	}

	s, ok := app["slots"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to get slot array")
	}

	slots := []*Slot{}

	for _, item := range s {
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		slot := &Slot{}
		if err := json.Unmarshal(b, slot); err != nil {
			return nil, err
		}
	}

	return slots, nil
}

func (client *Client) NewSlot(name, description, image string) (string, error) {

	slot := &Slot{
		AppID:       client.appID,
		SlotName:    name,
		Description: description,
		SlotImage:   image,
		//
		AcceptingCollections: true,
		IsPublic:             true,
	}

	data, err := client.Try("POST", "/api/v1/slot/new", slot)
	if err != nil {
		return "", err
	}
	m, ok := data.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to assert type: %v", data)
	}
	id := m["slotId"].(string)
	if !ok {
		return "", fmt.Errorf("failed to assert type: %v", m["slotId"])
	}
	return id, nil
}
