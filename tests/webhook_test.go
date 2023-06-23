package tests

import (
	"fmt"
	"testing"

	"github.com/pk-anderson/discord-messages/pkg"
)

func TestSendWithWebhook(t *testing.T) {
	message := "Test Message"

	//create client and service
	client := pkg.DiscordClient{
		WebHookEntity: pkg.WebHookEntity{
			URL: "https://discord.com/api/webhooks/1121519375174881310/QGYNsyc_qfNdY92Ax6Yj8NurCddQhg1GHkgA_y08FpYLMx9J3_-N4c-nmzdzMwN2XCqF",
		},
	}
	service := pkg.NewService(client)

	//test SendWithWebhook
	err := service.SendWithWebhook(message)
	if err != nil {
		t.Error("Erro ao enviar a mensagem:", err)
	}

	fmt.Println("Success")
}
