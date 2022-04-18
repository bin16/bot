package bot

import "fmt"

func (b *Bot) On(event string, fn UpdateHandler) {
	b.eventHandlers[event] = append(b.eventHandlers[event], fn)
	if b.Debug {
		for ev, hl := range b.eventHandlers {
			fmt.Printf("[ DEBUG ] - [ %s ], %d handlers\n", ev, len(hl))
		}
	}
}

func (b *Bot) Emit(event string, u *Update) {
	if b.Debug {
		fmt.Printf("[ DEBUG ] - Emit(%s, ...)\n", event)
	}

	hl := b.eventHandlers[event]
	for _, h := range hl {
		h(u, b)
	}
}
