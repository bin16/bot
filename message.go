package bot

const (
	MessageEntityTypeBotCommand = "bot_command"
)

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type Message struct {
	MessageID int64           `json:"message_id"`
	From      *User           `json:"from"`
	Chat      *Chat           `json:"chat"`
	Date      int64           `json:"date"`
	Text      string          `json:"text"`
	Animation *Animation      `json:"animation"`
	Entities  []MessageEntity `json:"entities"`
	Sticker   *Sticker        `json:"sticker"`
}
