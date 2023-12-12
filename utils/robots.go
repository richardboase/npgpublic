package utils

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func getRobotURL(target string, query *url.Values) (string, error) {

	host := "https://europe-west2-ninja-punk-girls.cloudfunctions.net/"

	switch target {
	case "prepare":
		host += "robot-prepare"

	default:
		return "", fmt.Errorf("no switch case found for %s", target)
	}

	return host + "?" + query.Encode(), nil
}

func CallRobot(target string, query *url.Values, dst interface{}) (interface{}, error) {

	host, err := getRobotURL(target, query)
	if err != nil {
		return nil, err
	}

	println("calling robot:", host)

	client := resty.New()

	response, err := client.R().SetResult(dst).Get(host)
	if err != nil {
		return nil, err
	}

	return response.Result(), nil
}
