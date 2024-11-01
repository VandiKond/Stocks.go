package messages

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Reply sends a message to the chat with an optional keyboard, splitting the text if it's too long.
func Reply(update tgbotapi.Update, text string, keyboard *tgbotapi.InlineKeyboardMarkup, Bot *tgbotapi.BotAPI) {
	// Check for valid text input
	if text == "" {
		return
	}

	// Split the message into parts if it exceeds 4096 characters
	messageParts := strings.Split(text, "\n")

	// Determine the ID of the message to reply to
	replyToMessageId := 0
	if update.Message != nil && update.Message.MessageID != 0 {
		replyToMessageId = update.Message.MessageID
	} else if update.Message != nil && update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.MessageID != 0 {
		replyToMessageId = update.Message.ReplyToMessage.MessageID
	}

	var sendText string

	// Send the message parts iteratively
	for i := 0; i < len(messageParts); i++ {
		// Checks the length of the text
		part := messageParts[i]
		leng := len(part) + len(sendText)
		ip1 := i + 1
		if leng >= 4000 || len(messageParts) == ip1 {
			// Gets the message
			message := update.Message
			if message == nil {
				message = update.CallbackQuery.Message
			}

			// Ads text
			if len(messageParts) == ip1 {
				sendText += "\n"
				sendText += part
			}

			// Creates the message
			msg := tgbotapi.NewMessage(message.Chat.ID, sendText)
			msg.ParseMode = "HTML"
			msg.ReplyToMessageID = replyToMessageId
			msg.DisableWebPagePreview = true
			if keyboard != nil && len(messageParts) == ip1 {
				msg.ReplyMarkup = keyboard
			}

			// Sends the message
			Bot.Send(msg)
			sendText = part
		} else {
			// Adding text
			sendText += "\n"
			sendText += part
		}
	}
}
