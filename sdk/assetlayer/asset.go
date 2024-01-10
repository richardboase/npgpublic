package assetlayer

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type Asset struct {
}

func (client *Client) MintAssets(collectionID string, quantity int) ([]string, error) {

	log.Println("minting collection:" + collectionID)

	data, err := client.Try(
		"POST",
		"/api/v1/asset/mint",
		nil,
		map[string]interface{}{
			"collectionId":    collectionID,
			"number":          quantity,
			"includeAssetIds": true,
		},
	)
	if err != nil {
		return nil, err
	}
	m, err := assertMapStringInterface(data)
	if err != nil {
		return nil, err
	}
	a, err := assertInterfaceArray(m["assetIds"])
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, item := range a {
		id, err := assertString(item)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (client *Client) AssetUpdate(assetID string, properties map[string]interface{}) error {

	_, err := client.Try(
		"PUT",
		"/api/v1/asset/update",
		nil,
		map[string]interface{}{
			"assetId":    assetID,
			"properties": properties,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) AssetExpressionValue(assetID, expressionID, attributeName string, properties map[string]interface{}, imageBytes []byte) error {

	var base64Encoding string
	// Determine the content type of the image file
	mimeType := http.DetectContentType(imageBytes)
	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	// Append the base64 encoded output
	base64Encoding += base64.StdEncoding.EncodeToString(imageBytes)

	_, err := client.Try(
		"POST",
		"/api/v1/asset/expressionValues",
		nil,
		map[string]interface{}{
			"assetId":                 assetID,
			"expressionAttributeName": attributeName,
			"expressionId":            expressionID,
			"value":                   base64Encoding,
		},
	)
	return err
}

func (client *Client) MintAssetWithProperties(collectionID string, object interface{}) (string, error) {
	ids, err := client.MintAssetsWithProperties(collectionID, object)
	if err != nil {
		return "", err
	}
	return ids[0], nil
}

func (client *Client) MintAssetsWithProperties(collectionID string, objects ...interface{}) ([]string, error) {
	list := []map[string]interface{}{}
	for _, object := range objects {
		b, err := json.Marshal(object)
		if err != nil {
			return nil, err
		}
		m := map[string]interface{}{}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	ids, err := client.MintAssets(collectionID, len(list))
	if err != nil {
		return nil, err
	}
	for n, id := range ids {
		if err := client.AssetUpdate(id, list[n]); err != nil {
			return nil, err
		}
	}
	return ids, nil
}

type Mint struct {
	Image      []byte
	Parameters interface{}
}

func (client *Client) MintAssetsWithPropertiesAndImage(collectionID, expressionID string, mints ...*Mint) error {
	list := []map[string]interface{}{}
	for _, mint := range mints {
		b, err := json.Marshal(mint.Parameters)
		if err != nil {
			return err
		}
		m := map[string]interface{}{}
		if err := json.Unmarshal(b, &m); err != nil {
			return err
		}
		list = append(list, m)
	}
	ids, err := client.MintAssets(collectionID, len(list))
	if err != nil {
		return err
	}
	for n, id := range ids {
		if err := client.AssetUpdate(id, list[n]); err != nil {
			return err
		}
		if err := client.AssetExpressionValue(id, expressionID, "Image", list[n], mints[n].Image); err != nil {
			return err
		}
	}
	return nil
}
