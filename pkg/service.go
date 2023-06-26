package pkg

type DiscordService interface {
	SendWithWebhook(message string) error
	SendWithBot(message string) error
}

type discordServiceImpl struct {
	DiscordClient DiscordClient
}

func (s *discordServiceImpl) SendWithWebhook(message string) error {
	//create client
	client := NewClient(s.DiscordClient)
	//send message
	err := client.SendWithWebhook(message)
	if err != nil {
		return err
	}
	//success
	return nil
}

func (s *discordServiceImpl) SendWithBot(message string) error {
	//create client
	client := NewClient(s.DiscordClient)
	//send message
	err := client.SendWithBot(message)
	if err != nil {
		return err
	}
	//success
	return nil
}

func NewService(DiscordClient DiscordClient) DiscordService {
	return &discordServiceImpl{
		DiscordClient: DiscordClient,
	}
}
