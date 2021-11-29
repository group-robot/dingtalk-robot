package robot

// TextMessage 文本消息
type TextMessage struct {
	BaseMessage
	At
	// Content 消息内容
	Content string
}

// NewTextMessage create TextMessage
func NewTextMessage() *TextMessage {
	message := new(TextMessage)
	message.MsgType = TextMsgType
	message.AtMobiles = []string{}
	message.AtUserIds = []string{}
	return message
}

// SetContent set TextMessage.Content
func (m *TextMessage) SetContent(content string) *TextMessage {
	m.Content = content
	return m
}

func (m *TextMessage) ToMessageMap() map[string]interface{} {
	atMap := m.ToAtMap()
	textMap := map[string]interface{}{"content": m.Content}

	message := m.ToMsgTypeMap()
	message["at"] = atMap
	message["text"] = textMap
	return message
}
