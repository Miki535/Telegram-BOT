package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	botToken := "6669079500:AAFFkfdH4m1P4Wm5Im-Mc21x2bfOv44gLKs"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)

			keyboard := tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton("btn").WithText("Hello World!"),
					tu.KeyboardButton("Location").WithRequestLocation(),
				),
			)

			message := tu.Message(
				chatID,
				"Keyboard",
			).WithReplyMarkup(keyboard)

			_, _ = bot.SendMessage(message)

		}
	}
}
