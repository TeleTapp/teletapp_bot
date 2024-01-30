package make_shop_bot

import (
	"app/config"
	"app/internal/api"
	"context"

	tele "gopkg.in/telebot.v3"
)

func startHandler(ctx tele.Context) error {
	c := api.NewClient()
	res, err := c.BotAuth(context.Background(), ctx.Chat().ID)
	if err != nil {
		return err
	}

	println(res.AccessToken)

	webApp := tele.WebApp{
		URL: config.App.WebAppBaseURL + "/bot/create",
	}

	btn := tele.MenuButton{
		Text:   "Open",
		Type:   tele.MenuButtonWebApp,
		WebApp: &webApp,
	}

	if err := ctx.Bot().SetMenuButton(ctx.Sender(), &btn); err != nil {
		return err
	}

	menu := &tele.ReplyMarkup{}
	menu.Inline(
		menu.Row(menu.WebApp("Open", &webApp)),
	)

	return ctx.Send("привет!", menu)
}
