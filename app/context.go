package app

import (
	"encoding/json"
	"github.com/ddliu/webhook/context"
	"io/ioutil"
	"net/http"
)

func buildContextFromRequest(c *context.Context, req *http.Request) error {
	var reqData map[string]interface{}
	// Method
	reqData["method"] = req.Method

	reqData["proto"] = req.Proto

	reqData["proto_major"] = req.ProtoMajor

	reqData["proto_minor"] = req.ProtoMinor

	// Header
	headers := make(map[string]string)
	for name, value := range req.Header {
		headers[name] = value[0]
	}
	reqData["headers"] = headers

	// payload
	if req.Header.Get("Content-Type") == "application/json" {
		var payload interface{}
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, payload)
		if err != nil {
			return err
		}

		reqData["payload"] = payload
	} else {
		// Form
		var form map[string]interface{}
		err := req.ParseForm()
		if err != nil {
			return err
		}

		for name, value := range req.PostForm {
			form[name] = value[0]
		}

		reqData["form"] = form
	}

	// Query
	query := make(map[string]string)
	for name, value := range req.URL.Query() {
		query[name] = value[0]
	}

	reqData["query"] = query

	c.SetValue("req", reqData)
	return nil
}