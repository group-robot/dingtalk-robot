package robot

type MsgType string

const (
	TextMsgType       MsgType = "text"
	LinkMsgType       MsgType = "link"
	MarkdownMsgType   MsgType = "markdown"
	ActionCardMsgType MsgType = "actionCard"
	FeedCardMsgType   MsgType = "feedCard"
)

type Message interface {
	ToMessageMap() map[string]interface{}
}
type BaseMessage struct {
	MsgType MsgType
}

func (m *BaseMessage) ToMsgTypeMap() map[string]interface{} {
	return map[string]interface{}{"msgtype": m.MsgType}
}

// At @
type At struct {
	// AtMobiles 被@人的手机号
	AtMobiles []string
	// AtUserIds 被@人的用户userid
	AtUserIds []string
	// IsAtAll 是否@所有人
	IsAtAll bool
}

func (m *At) ToAtMap() map[string]interface{} {
	atMap := map[string]interface{}{}
	if len(m.AtMobiles) > 0 {
		atMap["atMobiles"] = m.AtMobiles
	}
	if len(m.AtUserIds) > 0 {
		atMap["atUserIds"] = m.AtMobiles
	}
	atMap["isAtAll"] = m.IsAtAll
	return atMap
}

// SetAtMobiles set At.AtMobiles
func (m *At) SetAtMobiles(atMobiles []string) *At {
	m.AtMobiles = atMobiles
	return m
}

// AddAtMobiles add At.AtMobiles
func (m *At) AddAtMobiles(atMobiles ...string) *At {
	m.AtMobiles = append(m.AtMobiles, atMobiles...)
	return m
}

// SetAtUserIds set At.AtUserIds
func (m *At) SetAtUserIds(atUserIds []string) *At {
	m.AtUserIds = atUserIds
	return m
}

// AddAtUserIds add At.AtUserIds
func (m *At) AddAtUserIds(atUserIds ...string) *At {
	m.AtUserIds = append(m.AtUserIds, atUserIds...)
	return m
}

// SetIsAtAll set At.IsAtAll
func (m *At) SetIsAtAll(isAtAll bool) *At {
	m.IsAtAll = isAtAll
	return m
}

// AtAll set At.IsAtAll is true
func (m *At) AtAll() *At {
	m.IsAtAll = true
	return m
}
