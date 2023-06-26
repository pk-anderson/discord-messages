package pkg

import (
	"fmt"
	"net/http"
	"strings"
)

type Client interface {
	SendWithWebhook(message string) error
	SendWithBot(message string) error
}

type clientImpl struct {
	DiscordClient DiscordClient
}

func (s *clientImpl) SendWithWebhook(message string) error {
	//format message body
	body := fmt.Sprintf(`{"content":"%s"}`, strings.ReplaceAll(message, `"`, `\"`))

	//create request
	req, err := http.NewRequest("POST", s.DiscordClient.URL, strings.NewReader(body))
	if err != nil {
		return err
	}

	//set header
	req.Header.Set("Content-Type", "application/json")

	//create client and make request
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	//if status is not ok, error
	if err != nil {
		return err
	}

	//success
	return nil
}

func (s *clientImpl) SendWithBot(message string) error {
	//format url and message body
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", s.DiscordClient.ChannelId)
	body := fmt.Sprintf(`{"content":"%s"}`, strings.ReplaceAll(message, `"`, `\"`))

	//create request
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return err
	}

	//set header
	req.Header.Set("Authorization", "Bot "+s.DiscordClient.Token)
	req.Header.Set("Content-Type", "application/json")

	//create client and make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//if status is not ok, error
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send Discord message. Status: %s", resp.Status)
	}

	//success
	return nil
}

func NewClient(DiscordClient DiscordClient) Client {
	return &clientImpl{
		DiscordClient: DiscordClient,
	}
}
