package robot

// Link link
type Link struct {
	// Title 消息标题
	Title string
	// PicUrl 图片URL
	PicUrl string
	// MessageUrl 点击消息跳转的URL
	MessageUrl string
}

// NewLink create Link
func NewLink() *Link {
	return new(Link)
}

// SetTitle set Link.Title
func (m *Link) SetTitle(title string) *Link {
	m.Title = title
	return m
}

// SetPicUrl set Link.PicUrl
func (m *Link) SetPicUrl(picUrl string) *Link {
	m.PicUrl = picUrl
	return m
}

// SetMessageUrl set Link.MessageUrl
func (m *Link) SetMessageUrl(messageUrl string) *Link {
	m.MessageUrl = messageUrl
	return m
}
func (m *Link) toMap() map[string]interface{} {
	message := map[string]interface{}{}
	message["title"] = m.Title
	message["messageURL"] = m.MessageUrl
	message["picURL"] = m.PicUrl
	return message
}

// LinkMessage 链接消息类型
type LinkMessage struct {
	BaseMessage
	// Text  消息内容。如果太长只会部分展示
	Text string
	Link
}

// NewLinkMessage create LinkMessage
func NewLinkMessage() *LinkMessage {
	m := new(LinkMessage)
	m.MsgType = LinkMsgType
	return m
}

// SetTitle set  LinkMessage#Link.Title
func (m *LinkMessage) SetTitle(title string) *LinkMessage {
	m.Title = title
	return m
}

// SetText set LinkMessage#Links.text
func (m *LinkMessage) SetText(text string) *LinkMessage {
	m.Text = text
	return m
}

// SetPicUrl set LinkMessage#Link.PicUrl
func (m *LinkMessage) SetPicUrl(picUrl string) *LinkMessage {
	m.PicUrl = picUrl
	return m
}

// SetMessageUrl set LinkMessage#Link.MessageUrl
func (m *LinkMessage) SetMessageUrl(messageUrl string) *LinkMessage {
	m.MessageUrl = messageUrl
	return m
}

func (m *LinkMessage) ToMessageMap() map[string]interface{} {
	linkMap := map[string]interface{}{}
	linkMap["text"] = m.Text
	linkMap["title"] = m.Title
	linkMap["messageUrl"] = m.MessageUrl
	linkMap["picUrl"] = m.PicUrl

	message := m.ToMsgTypeMap()
	message["link"] = linkMap
	return message
}
