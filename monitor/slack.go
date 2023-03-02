package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SlackNotify struct {
	Username   string
	WebHookURL string
}

type SlackMessage struct {
	Username string `json:"username,omitempty"`
	Text     string `json:"text,omitempty"`
	Icon     string `json:"icon_emoji"`
}

func NewSlackNotify(username, webhook string) *SlackNotify {
	return &SlackNotify{
		Username:   username,
		WebHookURL: webhook,
	}
}

func (me SlackNotify) SendMessage(text string) error {
	if me.WebHookURL == "" {
		return nil
	}

	slackMsg, err := json.Marshal(
		&SlackMessage{
			Username: me.Username,
			Icon:     ":robot_face:",
			Text:     text,
		})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost,
		me.WebHookURL, bytes.NewBuffer(slackMsg))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack api is unavailable %d", resp.StatusCode)
	}

	return nil
}
