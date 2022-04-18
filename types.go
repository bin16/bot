package bot

import "encoding/json"

type APIResponse struct {
	OK          bool        `json:"ok"`
	Result      interface{} `json:"result"`
	Description string      `json:"description"`
	ErrorCode   int         `json:"error_code"`
}

func (r *APIResponse) Decode(target interface{}) error {
	return json.Unmarshal(r.Result.([]byte), target)
}

type User struct {
	ID       int64  `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Username string `json:"username"`
}

const (
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSuperGroup = "supergroup"
	ChatTypeChannel    = "channel"
)

type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Duration     int    `json:"duration"` // Duration of the video in seconds as defined by sender
	MIMEType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"` // Optional. File size in bytes
}

func (a *Animation) Ext() string {
	if a.MIMEType == "video/mp4" {
		return ".mp4"
	}

	return ".gif"
}

type UpdateConfig struct {
	Offset         int64    `json:"offset"` // getUpdates
	Limit          int      `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"allowed_updates"`
}
