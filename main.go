package main

import (
	"app/config"
	"app/internal/make_shop_bot"
	"app/logger"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	tele "gopkg.in/telebot.v3"
)

func main() {
	config.Load()
	l := logger.NewLogger("main")

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	if err := make_shop_bot.Init(close, &tele.Settings{
		Token:  config.App.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}); err != nil {
		l.Fatal().Err(err).Msg("failed to initialize make shop bot")
	}

	r := gin.Default()
	r.GET("/healthy", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	r.POST("/auth", func(ctx *gin.Context) {
		initData, _ := ioutil.ReadAll(ctx.Request.Body)

		fmt.Println(initdata.Validate(string(initData), config.App.Token, 24*time.Hour))

		ctx.JSON(200, gin.H{"kek": initData})
	})

	l.Info().Msg("service started")
	r.Run(config.App.Listen)
}
