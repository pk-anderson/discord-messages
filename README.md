### Lib para Envio de Mensagens para o Discord

Esta é uma biblioteca em Go (Golang) que fornece funcionalidades para enviar mensagens para o Discord usando webhooks ou a API de bot do Discord.

### Uso Básico

## Estrutura do Service 

```
type DiscordService interface {
	SendWithWebhook(message string) error
	SendWithBot(message string) error
}

type discordServiceImpl struct {
	DiscordClient DiscordClient
}

type DiscordClient struct {
	BotEntity
	WebHookEntity
}
```

As estruturas BotEntity e WebHookEntity são alias das estruturas presentes em internal/domain, escritas da seguinte forma: 

```
type WebHookEntity struct {
	URL string
}

type BotEntity struct {
	Token     string
	ChannelId string
}

```

## SendWithWebhook

A função SendWithWebhook é utilizada para enviar mensagens para um canal do discord, através de um Webhook criado para um canal específico. 

NECESSÁRIO PASSAR A URL DO WEBHOOK EM WEBHOOKENTITY NA INSTANCIAÇÃO DO DISCORDCLIENT.

Seu funcionamento pode ser visto em webhook_test.go, na pasta tests

PARA CRIAR UM WEBHOOK EM UM CANAL NO DISCORD:

    1 - Faça login na sua conta do Discord

    2 - No painel esquerdo, clique com o botão direito do mouse na categoria ou em um canal existente em que deseja criar o canal. Em seguida, clique em "Create Channel" (Criar Canal).

    3 - Escolha um nome para o canal e, opcionalmente, defina uma descrição. Selecione as permissões desejadas para os usuários que acessarão o canal. Depois de configurar as opções, clique em "Create" (Criar).

    4 - Agora que você criou o canal de mensagens, você pode criar um webhook para ele. Para fazer isso, clique com o botão direito do mouse no canal que você criou, escolha "Edit Channel" (Editar Canal) e navegue até a guia "Integrations" (Integrações) no menu lateral esquerdo.

    5 - Na guia "Integrations", clique em "Webhooks" e depois em "New Webhook" (Novo Webhook).

    6 - Forneça um nome para o webhook e, opcionalmente, escolha um avatar para representá-lo no canal. Em seguida, clique em "Copy Webhook URL" (Copiar URL do Webhook) para copiar o URL do webhook para a área de transferência.

## SendWithBot

A função SendWithBot é utilizada para enviar mensagens para um canal do discord, através de um Bot criado em um aplicativo dentro de uma conta no discord, sendo que este bot faz requisições para a api.

NECESSÁRIO PASSAR O TOKEN DO BOT E O ID DO CANAL DO DISCORD NA ENTIDADE BOTENTITY NA INSTANCIAÇÃO DO DISCORDCLIENT.

Seu funcionamento pode ser visto em bot_test.go, na pasta tests

PARA CRIAR BOT NO DISCORD E ADICIONAR PERMISSÕES:

    1 - Acesse o site do Discord Developer Portal em https://discord.com/developers/applications.

    2 - Faça login na sua conta do Discord, caso ainda não esteja logado.

    3 - Clique no botão "New Application" (Nova Aplicação) no canto superior direito.

    4 - Dê um nome para a sua aplicação (que será o nome do seu bot) e clique em "Create" (Criar).

    5 - Na página da sua nova aplicação, clique na guia "Bot" no menu lateral esquerdo.

    6 - Clique em "Add Bot" (Adicionar Bot) e confirme a ação clicando em "Yes, do it!" (Sim, faça isso!).

    7 - Agora você verá as informações do seu bot, incluindo o "TOKEN". Clique em "Copy" (Copiar) para copiar o token do seu bot. Guarde-o em um local seguro, pois é uma credencial importante.

ADICIONANDO BOT AO SERVIDOR: 

    1 - Acesse o Discord Developer Portal em https://discord.com/developers/applications.

    2 - Faça login na sua conta do Discord, se ainda não estiver logado.

    3 - Clique na aplicação correspondente ao seu bot.

    4 - No menu lateral esquerdo, clique na guia "OAuth2".

    5 - Em "Scopes" (Escopos), marque a opção "bot".

    6 - Logo abaixo, serão exibidos os escopos adicionais. Neste caso, marque a opção "Send Messages" (Enviar Mensagens) para conceder ao bot a permissão para enviar mensagens.

    7 - Na seção "Bot Permissions" (Permissões do Bot), selecione todas as permissões que você deseja conceder ao seu bot. Para enviar mensagens, você precisará pelo menos da permissão "Send Messages" (Enviar Mensagens).

    8 - Após selecionar as permissões desejadas, você verá um URL gerado em "Scopes". Copie esse URL.

    9 - Cole o URL em um navegador da web e selecione um servidor no qual deseja adicionar o seu bot.

CONCEDER PERMISSÕES AO BOT: 

    1 - No Discord, clique com o botão direito do mouse no nome do canal desejado e selecione "Edit Channel" (Editar Canal).

    2 - Na janela de configurações do canal, vá para a guia "Permissions" (Permissões).

    3 - Na seção "Roles/Members" (Funções/Membros), clique em "+ Add Role" (+ Adicionar Função) para adicionar uma nova função.

    4 - Crie uma nova função e defina as permissões necessárias para essa função. Certifique-se de habilitar a permissão "Send Messages" (Enviar Mensagens) para permitir que o bot envie mensagens no canal.

    5 - Depois de criar a nova função com as permissões desejadas, clique em "Save Changes" (Salvar Alterações) para aplicar as configurações.

    6 - Agora, vá para a seção "Roles/Members" (Funções/Membros) novamente e encontre a entrada do seu bot na lista.

    7 - Clique na caixa ao lado da nova função que você criou para atribuí-la ao bot.

Para obter o Id do canal:

    1 - Certifique-se de ter permissões suficientes na sua conta Discord para visualizar IDs (necessário ativar Modo Desenvolvedor).

    2 - No Discord, vá até o canal desejado.

    3 - Clique com o botão direito do mouse no nome do canal no painel esquerdo.

    4 - No menu suspenso, clique em "Copy ID" (Copiar ID). Se você não vê essa opção, verifique se você ativou o modo desenvolvedor nas configurações do Discord.

    5 - O ID do canal será copiado para a área de transferência. Agora você pode colar em algum lugar para usar posteriormente.