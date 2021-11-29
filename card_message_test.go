package robot

import (
	"os"
	"testing"
)

func TestNewSingleActionCardMessage(t *testing.T) {
	c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
	m := NewSingleActionCardMessage()
	m.SetTitle("乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身")
	m.SetText("![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png) \n ### 乔布斯 20 年前想打造的苹果咖啡厅 \n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划")
	m.SetBtnOrientation(Longitudinal)
	m.SetSingleTitle("阅读全文")
	m.SetSingleUrl("https://www.dingtalk.com/")
	_, err := c.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewActionCardMessage(t *testing.T) {
	c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
	m := NewActionCardMessage()
	m.SetTitle("乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身")
	m.SetText("![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png) \n ### 乔布斯 20 年前想打造的苹果咖啡厅 \n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划")
	m.SetBtnOrientation(Longitudinal)

	m.AddBtns(NewActionCardButton().SetTitle("内容不错").SetActionUrl("https://www.dingtalk.com/"), NewActionCardButton().SetTitle("不感兴趣").SetActionUrl("https://www.dingtalk.com/"))
	_, err := c.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewFeedCardMessage(t *testing.T) {
	c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
	m := NewFeedCardMessage()
	m.AddLinks(
		NewLink().SetTitle("时代的火车向前开1").SetMessageUrl("https://www.dingtalk.com/").SetPicUrl("https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png"),
		NewLink().SetTitle("时代的火车向前开2").SetMessageUrl("https://www.dingtalk.com/").SetPicUrl("https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png"),
	)
	_, err := c.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
