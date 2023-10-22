package log

import (
	"github.com/imroc/req/v3"
	"github.com/rotisserie/eris"
)

type EmbedWebhookMessage struct {
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
	Desc  string `json:"description,omitempty"`
	Color string `json:"color,omitempty"`
}

// IncomingWebhookMessage merupakan body request untuk incoming webhook discord, detail
// https://birdie0.github.io/discord-webhooks-guide/discord_webhook.html
type IncomingWebhookMessage struct {
	Content   string                `json:"content,omitempty"`
	Username  string                `json:"username,omitempty"`
	AvatarUrl string                `json:"avatar_url,omitempty"`
	TTS       bool                  `json:"tts,omitempty"`
	Embeds    []EmbedWebhookMessage `json:"embeds,omitempty"`
}

type discordWebhook struct {
	url string
}

// NewDiscordWebhook membuat objek discord webhook baru
func NewDiscordWebhook(url string) *discordWebhook {
	return &discordWebhook{
		url: url,
	}
}

// SendMessage menirim pesan ke server discord dengan tipe incoming webhook
func (dw *discordWebhook) SendMessage(m IncomingWebhookMessage) {
	client := req.C()

	resp, err := client.R().
		SetBody(m).
		Post(dw.url)

	if err != nil {
		ConsoleError(err)
	}

	if !resp.IsSuccessState() {
		ConsoleError(eris.New("Gagal mengirim log ke discord"))
	}
}
