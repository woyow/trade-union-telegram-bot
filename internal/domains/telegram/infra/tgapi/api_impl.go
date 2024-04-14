package tgapi

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"
)

const (
	chatIDLoggingKey     = "chatID"
	apiLoggingKey        = "api"
	telegramLoggingValue = "telegram"
)

type ApiImpl struct {
	api *echotron.API
	log *logrus.Logger
}

func NewApiImpl(api *echotron.API, log *logrus.Logger) *ApiImpl {
	return &ApiImpl{
		api: api,
		log: log,
	}
}

func (a *ApiImpl) SendMessage(dto entity.SendMessageApiDTO) error {
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
			apiLoggingKey:    telegramLoggingValue,
		}).Error("SendMessage - s.api.SendMessage error", err.Error())

		return err
	}

	if !resp.Ok {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			apiLoggingKey:    telegramLoggingValue,
		}).Error("SendMessage - status code: ", resp.ErrorCode)

		return errs.ErrStatusCodeUnsuccessful
	}

	return nil
}

func (a *ApiImpl) SendMessageWithoutError(dto entity.SendMessageApiDTO) {
	if err := a.SendMessage(dto); err != nil {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			apiLoggingKey:    telegramLoggingValue,
		}).Error("SendMessageWithoutError - a.SendMessage error", err.Error())
	}
}

func (a *ApiImpl) getInlineKeyboard(buttons [][]entity.InlineButton) [][]echotron.InlineKeyboardButton {
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

func (a *ApiImpl) SendMessageWithInlineKeyboard(dto entity.SendMessageWithInlineKeyboardApiDTO) error {
	inlineKeyboard := a.getInlineKeyboard(dto.Buttons)

	resp, err := a.api.SendMessage(dto.Text, dto.ChatID, &echotron.MessageOptions{
		ReplyMarkup: echotron.InlineKeyboardMarkup{
			InlineKeyboard: inlineKeyboard,
		},
	})
	if err != nil {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			apiLoggingKey:    telegramLoggingValue,
		}).Error("SendMessageWithInlineKeyboard - s.api.SendMessage error", err.Error())

		return err
	}

	if !resp.Ok {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			apiLoggingKey:    telegramLoggingValue,
		}).Error("SendMessageWithInlineKeyboard - status code: ", resp.ErrorCode)

		return errs.ErrStatusCodeUnsuccessful
	}

	return nil
}

func (a *ApiImpl) SendMessageWithInlineKeyboardWithoutError(dto entity.SendMessageWithInlineKeyboardApiDTO) {
	if err := a.SendMessageWithInlineKeyboard(dto); err != nil {
		a.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			apiLoggingKey:    telegramLoggingValue,
		}).Error("SendMessageWithInlineKeyboardWithoutError - a.SendMessageWithInlineKeyboard error", err.Error())
	}
}
