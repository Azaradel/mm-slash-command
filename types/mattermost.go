package types

import (
	"net/url"
)

type SlashCommandForm struct {
	ChannelID   string
	ChannelName string
	Command     string
	ResponseURL string
	TeamDomain  string
	TeamID      string
	Text        string
	Token       string
	UserID      string
	Username    string
}

func NewSlashCommandForm(form url.Values) *SlashCommandForm {
	return &SlashCommandForm{
		ChannelID:   form.Get("channel_id"),
		ChannelName: form.Get("channel_name"),
		Command:     form.Get("command"),
		ResponseURL: form.Get("response_url"),
		TeamDomain:  form.Get("team_domain"),
		TeamID:      form.Get("team_id"),
		Text:        form.Get("text"),
		Token:       form.Get("token"),
		UserID:      form.Get("user_id"),
		Username:    form.Get("username"),
	}
}
