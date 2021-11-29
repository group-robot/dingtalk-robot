package robot

type BtnOrientationType string

const (
	// Longitudinal 按钮竖直排列
	Longitudinal BtnOrientationType = "0"
	// Transverse 按钮横向排列
	Transverse BtnOrientationType = "1"
)

// ActionCard 卡片
type ActionCard struct {
	// Title 首屏会话透出的展示内容
	Title string
	// Text markdown格式的消息
	Text string
	// BtnOrientation btn 排列方式
	BtnOrientation BtnOrientationType
}

// SetTitle set ActionCard.Title
func (m *ActionCard) SetTitle(title string) *ActionCard {
	m.Title = title
	return m
}

// SetText set ActionCard.Text
func (m *ActionCard) SetText(text string) *ActionCard {
	m.Text = text
	return m
}

// SetBtnOrientation set ActionCard.BtnOrientation
func (m *ActionCard) SetBtnOrientation(btnOrientation BtnOrientationType) *ActionCard {
	m.BtnOrientation = btnOrientation
	return m
}
func (m ActionCard) toActionMap() map[string]interface{} {
	actionMap := map[string]interface{}{}
	actionMap["title"] = m.Title
	actionMap["text"] = m.Text
	actionMap["btnOrientation"] = m.BtnOrientation
	return actionMap
}

// ActionCardButton 跳转卡片中的按钮实体类
type ActionCardButton struct {
	// Title 按钮标题
	Title string
	// ActionUrl 点击按钮触发的URL
	ActionUrl string
}

// NewActionCardButton create ActionCardButton
func NewActionCardButton() *ActionCardButton {
	return &ActionCardButton{}
}

// SetTitle set ActionCardButton.Title
func (m *ActionCardButton) SetTitle(title string) *ActionCardButton {
	m.Title = title
	return m
}

// SetActionUrl set ActionCardButton.ActionUrl
func (m *ActionCardButton) SetActionUrl(actionUrl string) *ActionCardButton {
	m.ActionUrl = actionUrl
	return m
}
func (m *ActionCardButton) toActionMap() map[string]interface{} {
	actionMap := map[string]interface{}{}
	actionMap["title"] = m.Title
	actionMap["actionURL"] = m.ActionUrl
	return actionMap
}

// SingleActionCardMessage 整体跳转ActionCard类型
type SingleActionCardMessage struct {
	BaseMessage
	ActionCard
	// SingleTitle 单个按钮的标题
	SingleTitle string
	// 点击singleTitle按钮触发的URL
	SingleUrl string
}

// NewSingleActionCardMessage create SingleActionCardMessage
func NewSingleActionCardMessage() *SingleActionCardMessage {
	m := new(SingleActionCardMessage)
	m.MsgType = ActionCardMsgType
	return m
}

// SetSingleTitle set SingleActionCardMessage.SingleTitle
func (m *SingleActionCardMessage) SetSingleTitle(title string) *SingleActionCardMessage {
	m.SingleTitle = title
	return m
}

// SetSingleUrl set SingleActionCardMessage.SingleUrl
func (m *SingleActionCardMessage) SetSingleUrl(singleUrl string) *SingleActionCardMessage {
	m.SingleUrl = singleUrl
	return m
}

func (m *SingleActionCardMessage) ToMessageMap() map[string]interface{} {
	actionMap := m.toActionMap()
	actionMap["singleTitle"] = m.SingleTitle
	actionMap["singleURL"] = m.SingleUrl

	message := m.ToMsgTypeMap()
	message["actionCard"] = actionMap
	return message
}

// ActionCardMessage 独立跳转ActionCard类型
type ActionCardMessage struct {
	BaseMessage
	ActionCard
	Btns []*ActionCardButton
}

// NewActionCardMessage crate ActionCardMessage
func NewActionCardMessage() *ActionCardMessage {
	m := new(ActionCardMessage)
	m.MsgType = ActionCardMsgType
	m.Btns = []*ActionCardButton{}
	return m
}

// SetBtns set ActionCardMessage.Btns
func (m *ActionCardMessage) SetBtns(btns []*ActionCardButton) *ActionCardMessage {
	m.Btns = btns
	return m
}

// AddBtn add ActionCardMessage.Btns
func (m *ActionCardMessage) AddBtn(btn *ActionCardButton) *ActionCardMessage {
	m.Btns = append(m.Btns, btn)
	return m
}

// AddBtns add ActionCardMessage.Btns
func (m *ActionCardMessage) AddBtns(btn ...*ActionCardButton) *ActionCardMessage {
	m.Btns = append(m.Btns, btn...)
	return m
}

func (m *ActionCardMessage) ToMessageMap() map[string]interface{} {
	actionMap := m.toActionMap()
	btns := make([]map[string]interface{}, 0)
	for _, btn := range m.Btns {
		btns = append(btns, btn.toActionMap())
	}
	actionMap["btns"] = btns

	message := m.ToMsgTypeMap()
	message["actionCard"] = actionMap
	return message
}

// FeedCardMessage FeedCard类型
type FeedCardMessage struct {
	BaseMessage
	// Links 跳转链接
	Links []*Link
}

// NewFeedCardMessage create FeedCardMessage
func NewFeedCardMessage() *FeedCardMessage {
	m := new(FeedCardMessage)
	m.MsgType = FeedCardMsgType
	m.Links = []*Link{}
	return m
}

// SetLinks set FeedCardMessage.Links
func (m *FeedCardMessage) SetLinks(links []*Link) *FeedCardMessage {
	m.Links = links
	return m
}

// AddLinks add FeedCardMessage.Links
func (m *FeedCardMessage) AddLinks(links ...*Link) *FeedCardMessage {
	m.Links = append(m.Links, links...)
	return m
}
func (m *FeedCardMessage) ToMessageMap() map[string]interface{} {
	links := make([]map[string]interface{}, 0)
	for _, link := range m.Links {
		links = append(links, link.toMap())
	}
	linkMap := map[string]interface{}{}
	linkMap["links"] = links

	message := m.ToMsgTypeMap()
	message["feedCard"] = linkMap
	return message
}
