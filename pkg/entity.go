package pkg

type WebHookEntity struct {
	URL string
}

type BotEntity struct {
	Token     string
	ChannelId string
}

type DiscordClient struct {
	BotEntity
	WebHookEntity
}
