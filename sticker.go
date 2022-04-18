package bot

import (
	"fmt"
	"net/http"
	"strconv"
)

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type Sticker struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	IsAnimated   bool       `json:"is_animated"`
	IsVideo      bool       `json:"is_video"`
	Thumb        *PhotoSize `json:"thumb"`
	Emoji        string     `json:"emoji"`
	SetName      string     `json:"set_name"`
	FileSize     int        `json:"file_size"`
	// mask_position
}

func (s *Sticker) Ext() string {
	if s.IsAnimated {
		return ".tgs"
	}

	if s.IsVideo {
		return ".webm"
	}

	return ".png"
}

type CreateNewStickerSetConfig struct {
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Emojis      string    `json:"emojis"`
	PNGSticker  InputFile `json:"png_sticker" empty:"-"`
	WebMSticker FilePath  `json:"webm_sticker" empty:"-"`
	TGSSticker  FilePath  `json:"tgs_sticker" empty:"-"`
}

type AddStickerToSetConfig struct {
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Emojis      string    `json:"emojis"`
	PNGSticker  InputFile `json:"png_sticker" empty:"-"`
	WebMSticker FilePath  `json:"webm_sticker" empty:"-"`
	TGSSticker  FilePath  `json:"tgs_sticker" empty:"-"`
}

func (b *Bot) AddStickerToSet(d *AddStickerToSetConfig) error {
	form := NewFormData()
	form.Append("user_id", strconv.Itoa(int(d.UserID)))
	form.Append("emojis", d.Emojis)
	form.Append("name", d.Name)
	if d.WebMSticker != "" {
		form.AppendFile("webm_sticker", string(d.WebMSticker))
	} else if d.TGSSticker != "" {
		form.AppendFile("tgs_sticker", string(d.TGSSticker))
	}

	url := b.getAPI(MethodAddStickerToSet)
	resp, err := http.Post(url, form.ContentType(), form.Done())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := &APIResponse{Result: false}
	if err := parseResponse(resp.Body, r); err != nil {
		return err
	}

	fmt.Println(r)
	if !r.OK {
		return fmt.Errorf(r.Description)
	}

	return nil
}

func (b *Bot) CreateNewStickerSet(d *CreateNewStickerSetConfig) error {
	form := NewFormData()
	form.Append("user_id", strconv.Itoa(int(d.UserID)))
	form.Append("emojis", d.Emojis)
	form.Append("title", d.Title)
	form.Append("name", d.Name)
	if d.WebMSticker != "" {
		form.AppendFile("webm_sticker", string(d.WebMSticker))
	} else if d.TGSSticker != "" {
		form.AppendFile("tgs_sticker", string(d.TGSSticker))
	}

	url := b.getAPI(MethodCreateNewStickerSet)
	resp, err := http.Post(url, form.ContentType(), form.Done())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := &APIResponse{Result: false}
	if err := parseResponse(resp.Body, r); err != nil {
		return err
	}

	fmt.Println(r)
	if !r.OK {
		return fmt.Errorf(r.Description)
	}

	return nil
}
