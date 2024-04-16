package tgapi

import (
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"

	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
)

func (a *APIImpl) SendMessage(dto entity.SendMessageAPIDTO) error {
	options := &echotron.MessageOptions{}

	if dto.Options != nil {
		if dto.Options.ParseMode != nil {
			options.ParseMode = echotron.ParseMode(*dto.Options.ParseMode)
		}
	} else {
		options = nil
	}

	resp, err := a.api.SendMessage(dto.Text, dto.ChatID, options)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			infraLoggingKey:  infraLoggingValue,
			domainLoggingKey: domainLoggingValue,
			textLoggingKey:   dto.Text,
		}).Error("SendMessage - s.api.SendMessage error: ", err.Error())

		return err
	}

	if !resp.Ok {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			infraLoggingKey:  infraLoggingValue,
			domainLoggingKey: domainLoggingValue,
			textLoggingKey:   dto.Text,
		}).Error("SendMessage - status code: ", resp.ErrorCode)

		return errs.ErrStatusCodeUnsuccessful
	}

	return nil
}

func (a *APIImpl) getInlineKeyboard(buttons [][]entity.InlineButton) [][]echotron.InlineKeyboardButton {
	inlineKeyboard := make([][]echotron.InlineKeyboardButton, 0, len(buttons))

	for i := range buttons {
		inlineButtons := make([]echotron.InlineKeyboardButton, 0, len(buttons[i]))

		for j := range buttons[i] {
			inlineButtons = append(inlineButtons, echotron.InlineKeyboardButton{
				CallbackGame:                 nil,
				WebApp:                       nil,
				LoginURL:                     nil,
				SwitchInlineQueryChosenChat:  nil,
				Text:                         buttons[i][j].Text,
				CallbackData:                 buttons[i][j].CallbackData,
				SwitchInlineQuery:            "",
				SwitchInlineQueryCurrentChat: "",
				URL:                          buttons[i][j].URL,
				Pay:                          false,
			})
		}

		inlineKeyboard = append(inlineKeyboard, inlineButtons)
	}

	return inlineKeyboard
}

func (a *APIImpl) SendMessageWithInlineKeyboard(dto entity.SendMessageWithInlineKeyboardAPIDTO) error {
	inlineKeyboard := a.getInlineKeyboard(dto.Buttons)

	resp, err := a.api.SendMessage(dto.Text, dto.ChatID, &echotron.MessageOptions{
		ReplyMarkup: echotron.InlineKeyboardMarkup{
			InlineKeyboard: inlineKeyboard,
		},
	})
	if err != nil {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			infraLoggingKey:  infraLoggingValue,
			domainLoggingKey: domainLoggingValue,
			textLoggingKey:   dto.Text,
		}).Error("SendMessageWithInlineKeyboard - s.api.SendMessage error", err.Error())

		return err
	}

	if !resp.Ok {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			infraLoggingKey:  infraLoggingValue,
			domainLoggingKey: domainLoggingValue,
			textLoggingKey:   dto.Text,
		}).Error("SendMessageWithInlineKeyboard - status code: ", resp.ErrorCode)

		return errs.ErrStatusCodeUnsuccessful
	}

	return nil
}
