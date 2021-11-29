package robot

import (
	"os"
	"testing"
)

func TestNewMarkdownMessage(t *testing.T) {
	c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
	m := NewMarkdownMessage()
	m.SetTitle("杭州天气")
	m.SetText("#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n")
	m.AddAtMobiles(os.Getenv("mobile"))
	_, err := c.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
