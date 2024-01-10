package assetlayer

import (
	"encoding/json"
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

func (client *Client) DeleteSlot(slotID string) error {
	_, err := client.Try(
		"POST",
		"/api/v1/slot/remove",
		map[string]string{
			"appId":  client.appID,
			"slotId": slotID,
		},
	)
	return err
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

	data, err := client.Try("POST", "/api/v1/slot/new", nil, slot)
	if err != nil {
		return "", err
	}
	m, err := assertMapStringInterface(data)
	if err != nil {
		return "", err
	}
	id, err := assertString(m["slotId"])
	if err != nil {
		return "", err
	}
	return id, nil
}

func (client *Client) GetSlots() ([]*Slot, error) {

	data, err := client.Try(
		"GET",
		"/api/v1/app/slots",
		map[string]string{
			"appId":  client.appID,
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
	app, err := assertMapStringInterface(m["app"])
	if err != nil {
		return nil, err
	}
	s, err := assertInterfaceArray(app["slots"])
	if err != nil {
		return nil, err
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
		slots = append(slots, slot)
	}

	return slots, nil
}
