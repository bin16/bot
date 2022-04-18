package bot

import (
	"log"
)

type UpdateHandler func(u *Update, b *Bot)

type Bot struct {
	token        string
	UpdateConfig *UpdateConfig
	Debug        bool
	Profile      *User

	eventHandlers  map[string][]UpdateHandler
	updateHandlers []UpdateHandler
}

func New(token string) *Bot {
	return &Bot{
		token: token,
		Debug: true,
		UpdateConfig: &UpdateConfig{
			Offset:         1,
			Timeout:        30,
			Limit:          50,
			AllowedUpdates: []string{"message"},
		},
		eventHandlers: make(map[string][]UpdateHandler),
	}
}

func (b *Bot) Reply(m *Message, text string) (*Message, error) {
	m0 := Sendable{
		ChatID:           m.Chat.ID,
		Text:             text,
		ReplyToMessageID: m.MessageID,
	}
	m1 := &Message{}

	_, err := b.postJSON(MethodSendMessage, m0, m1)
	if err != nil {
		log.Println(err)
	}
	return m1, err
}

type Sendable struct {
	ChatID           int64  `json:"chat_id"`
	Text             string `json:"text"`
	ReplyToMessageID int64  `json:"reply_to_message_id"`
}

func (b *Bot) getMe() {
	u := &User{}
	b.postJSON(MethodGetMe, nil, u)
	b.Profile = u
}
