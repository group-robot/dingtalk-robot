# dingtalk-robot
[![Go Reference](https://pkg.go.dev/badge/github.com/hb0730/dingtalk-robot.svg)](https://pkg.go.dev/github.com/hb0730/dingtalk-robot) 钉钉自定义机器人

# Example

### text
```go
c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
m := NewTextMessage()
m.SetContent("test").AtAll()
_, err := c.Send(m)
if err != nil {
    t.Fatal(err)
}
```

### link
```go
c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
m := NewLinkMessage()
m.SetText("这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林")
m.SetTitle("时代的火车向前开")
m.SetMessageUrl("https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI")
_, err := c.Send(m)
if err != nil {
    t.Fatal(err)
}
```

### markdown
```go
c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
m := NewMarkdownMessage()
m.SetTitle("杭州天气")
m.SetText("#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n")
m.AddAtMobiles(os.Getenv("mobile"))
_, err := c.Send(m)
if err != nil {
    t.Fatal(err)
}
```

### card
#### single action card
```go
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
```

#### action card
```go
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
```

#### feed card
```go
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
```