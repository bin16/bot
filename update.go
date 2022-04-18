package bot

import "log"

type Update struct {
	UpdateID          int64    `json:"update_id"`
	Message           *Message `json:"message"`             // Optional
	EditedMessage     *Message `json:"edited_message"`      // Optional
	ChannelPost       *Message `json:"channel_post"`        // Optional
	EditedChannelPost *Message `json:"edited_channel_post"` // Optional
	// inline_query
	// chosen_inline_result
	// ...
}

func (u *Update) Type() string {
	if u.Message != nil {
		return "message"
	}

	if u.EditedMessage != nil {
		return "edited_message"
	}

	if u.ChannelPost != nil {
		return "channel_post"
	}

	if u.EditedChannelPost != nil {
		return "edited_channel_post"
	}

	return ""
}

func (b *Bot) getUpdates(conf UpdateConfig) []Update {
	ul := []Update{}
	b.postJSON(MethodGetUpdates, conf, &ul)
	return ul
}

func (b *Bot) handleUpdate(u *Update) {
	hl := b.updateHandlers
	for _, h := range hl {
		h(u, b)
	}
}

func (b *Bot) Run() {
	b.getMe()

	for {
		ul := b.getUpdates(*b.UpdateConfig)
		if b.Debug {
			log.Printf("getUpdates() %d\n", len(ul))
		}
		for _, u := range ul {
			if u.UpdateID > b.UpdateConfig.Offset {
				b.UpdateConfig.Offset = u.UpdateID
			}

			b.handleUpdate(&u)
		}

		b.UpdateConfig.Offset++
	}
}

func (b *Bot) OnUpdate(fn UpdateHandler) {
	b.updateHandlers = append(b.updateHandlers, fn)
}
