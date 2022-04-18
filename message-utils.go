package bot

import (
	"fmt"
	"regexp"
)

func (m *Message) IsPrivate() bool {
	return m.Chat.Type == ChatTypePrivate
}

func (m *Message) IsGIF() bool {
	if m.Animation == nil {
		return false
	}

	return m.Animation.MIMEType == "video/mp4"
}

func (m *Message) IsText() bool {
	return m.Text != ""
}

func (m *Message) Command() string {
	rx := regexp.MustCompile(`/([\w]+)@?`)
	for _, t := range m.Entities {
		if t.Type == MessageEntityTypeBotCommand {
			s := m.Text[t.Offset:t.Length]
			fmt.Println(s)
			results := rx.FindStringSubmatch(s)
			return results[1]
		}
	}

	return ""
}
