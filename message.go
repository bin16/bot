package bot

import (
	"io"
	"reflect"
	"strconv"
)

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

type SendMessageConfig struct {
	ChatID           int64  `json:"chat_id"`
	Text             string `json:"text"`
	ReplyToMessageID int64  `json:"reply_to_message_id"`
}

type SendDocumentConfig struct {
	ChatID           int64     `json:"chat_id"`
	Document         InputFile `json:"document"`
	ReplyToMessageID int64     `json:"reply_to_message_id"`
	form             *FormData
}

func (d *SendDocumentConfig) Args(token string) (url, contentType string, body io.Reader) {
	d.form = NewFormData().Init()
	d.form.Append("chat_id", strconv.Itoa(int(d.ChatID)))
	if d.ReplyToMessageID != 0 {
		d.form.Append("reply_to_message_id", strconv.Itoa(int(d.ReplyToMessageID)))
	}

	if reflect.TypeOf(d.Document).Kind() == reflect.TypeOf(FilePath("")).Kind() {
		d.form.AppendFile("document", string(d.Document.(FilePath)))
	} else if reflect.TypeOf(d.Document).Kind() == reflect.TypeOf(FileID("")).Kind() {
		d.form.Append("document", string(d.Document.(FileID)))
	}

	return getAPI(token, MethodSendDocument), d.form.ContentType(), d.form.Done()
}
