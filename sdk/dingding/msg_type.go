package dingding

import "fmt"

const (
	MsgTypeTextTmpl       = `{"msgtype":"text","text":{"content":"%s"}}`
	MsgTypeMarkdownTmpl   = `{"msgtype":"markdown","markdown":{"title":"%s","text":"%s"}}`
	MsgTypeMarkdownAtTmpl = `{"msgtype":"markdown","markdown":{"title":"%s","text":"%s"},"at":{"isAtAll": false, "atMobiles":[]}}`
)

type RobotConfig struct {
	Token  string
	Secret string
}

type DingMsg struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
	At      At     `json:"at"`
}

type Text struct {
	Content string `json:"content"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []int    `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type ActionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
}

type FeedCard struct {
	Links []struct {
		Title      string `json:"title"`
		MessageURL string `json:"messageURL"`
		PicURL     string `json:"picURL"`
	} `json:"links"`
}
type MsgType int

const (
	_ MsgType = iota
	MsgTypeText
	MsgTypeMarkdown
	MsgTypeLink
	MsgTypeActionCard
	MsgTypeFeedCard
)

func (c MsgType) String() string {
	switch c {
	case MsgTypeText:
		return "text"
	case MsgTypeMarkdown:
		return "markdown"
	case MsgTypeLink:
		return "link"
	case MsgTypeActionCard:
		return "actionCard"
	case MsgTypeFeedCard:
		return "feedCard"
	default:
		return ""
	}
}

func (c MsgType) Tmpl() string {
	switch c {
	case MsgTypeText:
		return MsgTypeTextTmpl
	case MsgTypeMarkdown:
		return MsgTypeMarkdownTmpl
	default:
		return MsgTypeTextTmpl
	}
}

func (c MsgType) Body(title, content string) string {
	switch c {
	case MsgTypeText:
		return fmt.Sprintf(c.Tmpl(), content)
	case MsgTypeMarkdown:
		return fmt.Sprintf(c.Tmpl(), title, content)
	default:
		return fmt.Sprintf(c.Tmpl(), content)
	}
}
