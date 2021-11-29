package robot

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hb0730/dingtalk-robot/internal/security"
	"net/url"
	"strconv"
	"time"
)

// Client dingtalk robot client
type Client struct {
	WebHok string
	Secret string
}

// NewClient create Client
func NewClient(webHok, secret string) *Client {
	return &Client{WebHok: webHok, Secret: secret}
}

// ErrorResponse response error msg
type ErrorResponse struct {
	Code   int    `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func (c *Client) Send(message Message) (*ErrorResponse, error) {
	res := &ErrorResponse{}
	if len(c.WebHok) == 0 {
		return res, fmt.Errorf("wehok is blank")
	}
	if message == nil {
		return res, fmt.Errorf("body is null")
	}
	timestamp := time.Now().UnixNano() / 1e6
	sign, err := security.GenSign(c.Secret, timestamp)
	if err != nil {
		return res, err
	}
	client := resty.New()
	resp, err := client.SetRetryCount(3).R().
		SetQueryString(c.urlParamsEncode(sign, timestamp)).
		SetBody(message.ToMessageMap()).
		SetHeader("Accept", "application/json").
		SetHeader("Accept-Charset", "utf8").
		SetHeader("Content-Type", "application/json").
		SetResult(&ErrorResponse{}).
		ForceContentType("application/json").
		Post(c.WebHok)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*ErrorResponse)
	if result.Code != 0 {
		return nil, fmt.Errorf("send message to dingtalk robot error: %s", result.ErrMsg)
	}
	return result, nil
}
func (c *Client) urlParamsEncode(sign string, timestamp int64) string {
	value := url.Values{}
	value.Set("timestamp", strconv.FormatInt(timestamp, 10))
	value.Set("sign", sign)
	return value.Encode()
}
