package make_shop_bot

import (
	"os"

	tele "gopkg.in/telebot.v3"
)

func Init(close chan os.Signal, s *tele.Settings) error {
	b, err := tele.NewBot(*s)
	if err != nil {
		return err
	}

	b.Handle("/start", startHandler)

	go func() {
		defer b.Stop()
		b.Start()
		<-close
	}()

	return nil
}
