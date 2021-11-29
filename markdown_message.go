package robot

// MarkdownMessage markdown类型
type MarkdownMessage struct {
	BaseMessage
	At
	// Title  首屏会话透出的展示内容
	Title string
	// Text markdown格式的消息
	Text string
}

// NewMarkdownMessage create MarkdownMessage
func NewMarkdownMessage() *MarkdownMessage {
	m := new(MarkdownMessage)
	m.AtMobiles = []string{}
	m.AtUserIds = []string{}
	m.MsgType = MarkdownMsgType
	return m
}

// SetTitle set MarkdownMessage.Title
func (m *MarkdownMessage) SetTitle(title string) *MarkdownMessage {
	m.Title = title
	return m
}

// SetText set MarkdownMessage.Text
func (m *MarkdownMessage) SetText(text string) *MarkdownMessage {
	m.Text = text
	return m
}
func (m *MarkdownMessage) ToMessageMap() map[string]interface{} {
	markdownMap := map[string]interface{}{}
	markdownMap["title"] = m.Title
	markdownMap["text"] = m.Text
	atMap := m.ToAtMap()

	message := m.ToMsgTypeMap()
	message["markdown"] = markdownMap
	message["at"] = atMap
	return message
}
