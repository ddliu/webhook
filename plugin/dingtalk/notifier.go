package dingtalk

import (
	"encoding/json"
	"errors"
	"github.com/ddliu/go-httpclient"
	"github.com/ddliu/webhook/contact"
	"github.com/ddliu/webhook/context"
	"github.com/ddliu/webhook/notifier"
	"github.com/spf13/cast"
)

// DingRobot notifier
// See https://open-doc.dingtalk.com/docs/doc.htm?treeId=257&articleId=105735&docType=1
type DingRobot struct {
}

const DING_ROBOT_API_URL = "https://oapi.dingtalk.com/robot/send"

func (d *DingRobot) GetId() string {
	return "ding_robot"
}

func (d *DingRobot) Config(c *context.Context) {

}

type DingMessageText struct {
	MsgType string `json:"msgtype,"`
	Text    struct {
		Content string `json:"content,"`
	} `json:"text,"`
}

type DingResponse struct {
	ErrCode int    `json:"errcode,"`
	ErrMsg  string `json:"errmsg,"`
}

func (d *DingRobot) Notify(c *contact.Contact, title, content string) error {
	p := c.GetProperty("DingRobot")
	ps := cast.ToString(p)

	msg := DingMessageText{}
	msg.MsgType = "text"
	msg.Text.Content = content
	if title != "" {
		msg.Text.Content = title + "\n\n" + msg.Text.Content
	}

	resp, err := httpclient.PostJson(DING_ROBOT_API_URL+"?access_token="+ps, msg)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Error response")
	}

	body, err := resp.ReadAll()
	if err != nil {
		return err
	}

	var dingResponse DingResponse
	if err := json.Unmarshal(body, &dingResponse); err != nil {
		return err
	}

	if dingResponse.ErrCode != 0 {
		return errors.New(dingResponse.ErrMsg)
	}

	return nil
}

func (d *DingRobot) IsMatch(c *contact.Contact) bool {
	p := c.GetProperty("DingRobot")
	return p != nil
}

func init() {
	notifier.RegisterNotifier(&DingRobot{})
}
