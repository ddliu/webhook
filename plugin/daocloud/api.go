package daocloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ddliu/go-httpclient"
	log "github.com/sirupsen/logrus"
)

const DAOCLOUD_API_SERVER = "https://openapi.daocloud.io"

func callApi(method string, api string, token string, data interface{}, response interface{}) error {
	var body []byte
	var err error
	api = DAOCLOUD_API_SERVER + api
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	resp, err := httpclient.Do(method, api, map[string]string{
		"Authorization": token,
	}, bytes.NewReader(body))

	if err != nil {
		return err
	}

	respBody, err := resp.ReadAll()
	if err != nil {
		log.WithFields(log.Fields{
			"Plugin": "daocloud",
		}).Debugf("CallApi %s error: %s", api, err.Error())
		return err
	}

	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"Plugin": "daocloud",
			"Api":    api,
			"Token":  token,
			"Data":   data,
		}).Debugf("CallApi error code %d, body: ", resp.StatusCode, string(respBody))
		return errors.New("Response code not 200")
	}

	return json.Unmarshal(respBody, &data)
}
