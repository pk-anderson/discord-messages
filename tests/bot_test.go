package tests

import (
	"fmt"
	"testing"

	"github.com/pk-anderson/discord-messages/pkg"
)

func TestSendWithBot(t *testing.T) {
	message := "Test Message"

	//create client and service
	client := pkg.DiscordClient{
		BotEntity: pkg.BotEntity{
			Token:     "MTEyMTUyNDI4OTAyODUwNTc0MA.G8i3aj.SHvSEA8hykPHcqsXKRGfTULi6qlsXaf2WUuuAo",
			ChannelId: "1121519176033521725",
		},
	}
	service := pkg.NewService(client)

	//test SendWithBot
	err := service.SendWithBot(message)
	if err != nil {
		t.Error("Erro ao enviar a mensagem:", err)
	}

	fmt.Println("Success")
}
