package log

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/rotisserie/eris"
)

type EmbedWebhookMessage struct {
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
	Desc  string `json:"description,omitempty"`
	Color string `json:"color,omitempty"`
}

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
	body, _ := json.Marshal(m)
	postBody := bytes.NewBuffer(body)

	resp, err := http.Post(dw.url, "application/json", postBody)
	if err != nil {
		ConsoleError(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			ConsoleError(err)
		}

		ConsoleError(eris.New(string(b)))
	}

}
