package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"strconv"
	"strings"

	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"
)

func (s *Service) NewCommand(ctx context.Context, dto entity.NewCommandServiceDTO) error {
	if err := s.deleteDraftAppeal(ctx, deleteDraftAppealDTO{
		chatID: dto.ChatID,
	}); err != nil {
		s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
			lang:   dto.Lang,
			chatID: dto.ChatID,
		})

		return err
	}

	out, err := s.createAppeal(ctx, createAppealDTO{
		chatID:  dto.ChatID,
		isDraft: true,
	})
	if err != nil {
		s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
			lang:   dto.Lang,
			chatID: dto.ChatID,
		})

		return err
	}

	s.log.WithFields(logrus.Fields{
		chatIDLoggingKey: dto.ChatID,
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Info("NewCommand - Draft appeal has been created with id: ", out.ID)

	if err := s.api.SendMessage(entity.SendMessageAPIDTO{
		Text:    s.translate(newCommandTranslateKey, dto.Lang),
		ChatID:  dto.ChatID,
		Options: nil,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("NewCommand - s.api.SendMessage error: ", err.Error())
	}

	if err := s.api.SendMessage(entity.SendMessageAPIDTO{
		Text:    s.translate(newCommandTypeFullNameTranslateKey, dto.Lang),
		ChatID:  dto.ChatID,
		Options: nil,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("NewCommand - s.api.SendMessage error: ", err.Error())
	}

	return nil
}

type deleteDraftAppealDTO struct {
	chatID int64
}

func (s *Service) deleteDraftAppeal(ctx context.Context, dto deleteDraftAppealDTO) error {
	if err := s.repo.DeleteDraftAppeal(ctx, entity.DeleteDraftAppealRepoDTO{
		ChatID: dto.chatID,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("deleteDraftAppeal - s.repo.DeleteDraftAppeal error: ", err.Error())

		return err
	}

	return nil
}

type createAppealDTO struct {
	chatID  int64
	isDraft bool
}

func (s *Service) createAppeal(ctx context.Context, dto createAppealDTO) (*entity.CreateAppealOut, error) {
	out, err := s.repo.CreateAppeal(ctx, entity.CreateAppealRepoDTO{
		ChatID:  dto.chatID,
		IsDraft: dto.isDraft,
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("createAppeal - s.repo.CreateAppeal error: ", err.Error())
		return nil, err
	}

	return out, nil
}

type getButtonsDTO struct {
	buttons          []entity.InlineButton
	maxButtonsOnLine int
}

func (s *Service) getButtons(dto getButtonsDTO) [][]entity.InlineButton {
	lines := make([][]entity.InlineButton, 0, int(math.Ceil(float64(len(dto.buttons)/dto.maxButtonsOnLine))))

	for i := 0; i < len(dto.buttons); i++ {
		buttons := make([]entity.InlineButton, 0, dto.maxButtonsOnLine)

		buttonsOnLine := func() int {
			if len(dto.buttons)-i >= dto.maxButtonsOnLine {
				return dto.maxButtonsOnLine
			}

			return len(dto.buttons) - i
		}()

		for j := 0; j < buttonsOnLine; j++ {
			s.log.Debug("getButtons - dto.buttons[i].Text: ", dto.buttons[i].Text)

			buttons = append(buttons, entity.InlineButton{
				Text:         strconv.Itoa(i + 1),
				CallbackData: dto.buttons[i].CallbackData,
				URL:          dto.buttons[i].URL,
			})

			if j != buttonsOnLine-1 {
				i++
			}
		}

		lines = append(lines, buttons)
	}

	return lines
}

func (s *Service) NewCommandFullNameState(ctx context.Context, dto entity.NewCommandFullNameStateServiceDTO) error {
	if err := s.repo.UpdateDraftAppeal(ctx, entity.UpdateAppealRepoDTO{
		UpdateAppealBase: entity.UpdateAppealBase{
			Fname:   nil,
			Lname:   nil,
			Mname:   &dto.Text,
			Subject: nil,
			IsDraft: nil,
		},
		ChatID: dto.ChatID,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("NewCommandFullNameState - s.repo.UpdateDraftAppeal error: ", err.Error())

		s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
			lang:   dto.Lang,
			chatID: dto.ChatID,
		})

		return err
	}

	isActiveSubjects := true

	subjects, err := s.repo.GetAppealSubjects(ctx, entity.GetAppealSubjectsRepoDTO{
		ChatID:   dto.ChatID,
		IsActive: &isActiveSubjects,
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("NewCommandFullNameState - s.repo.GetAppealSubjects error: ", err.Error())

		s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
			lang:   dto.Lang,
			chatID: dto.ChatID,
		})

		return err
	}

	buttons := s.getButtons(getButtonsDTO{
		buttons: func() []entity.InlineButton {
			buttons := make([]entity.InlineButton, 0, len(subjects))

			for i := range subjects {
				buttons = append(buttons, entity.InlineButton{
					Text:         subjects[i].Text[dto.Lang],
					CallbackData: subjects[i].CallbackData,
					URL:          "",
				})
			}

			return buttons
		}(),
		maxButtonsOnLine: 5,
	})

	if err := s.api.SendMessageWithInlineKeyboard(entity.SendMessageWithInlineKeyboardAPIDTO{
		Text: s.translate(newCommandChooseSubjectOfAppealTranslateKey, dto.Lang) +
			"\n\n" +
			func() string {
				var text string
				for i := range subjects {
					text += fmt.Sprintf("%d. %s\n", i+1, s.translateLocalization(subjects[i].Text, dto.Lang))
				}
				return text
			}(),
		ChatID:  dto.ChatID,
		Buttons: buttons,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("NewCommandFullNameState - s.api.SendMessage error: ", err.Error())
	}

	return nil
}

func (s *Service) NewCommandSubjectState(ctx context.Context, dto entity.NewCommandSubjectStateServiceDTO) error {
	if err := s.repo.UpdateDraftAppeal(ctx, entity.UpdateAppealRepoDTO{
		UpdateAppealBase: entity.UpdateAppealBase{
			Fname:   nil,
			Lname:   nil,
			Mname:   nil,
			Subject: &dto.Data,
			IsDraft: nil,
		},
		ChatID: dto.ChatID,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("NewCommandLastNameState - s.repo.UpdateDraftAppeal error: ", err.Error())

		s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
			lang:   dto.Lang,
			chatID: dto.ChatID,
		})

		return err
	}

	out, err := s.repo.GetDraftAppeal(ctx, entity.GetDraftAppealRepoDTO{
		ChatID: dto.ChatID,
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("NewCommandSubjectState - s.repo.GetDraftAppeal error: ", err.Error())

		s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
			lang:   dto.Lang,
			chatID: dto.ChatID,
		})

		return err
	}

	_ = s.sendFinalAppeal(sendFinalAppealDTO{
		appeal: out,
		lang:   dto.Lang,
		chatID: dto.ChatID,
	})

	_ = s.sendConfirm(sendConfirmDTO{
		lang:   dto.Lang,
		chatID: dto.ChatID,
	})

	return nil
}

type sendFinalAppealDTO struct {
	appeal *entity.GetDraftAppealRepoOut
	lang   string
	chatID int64
}

func (s *Service) sendFinalAppeal(dto sendFinalAppealDTO) error {
	var text string

	fname := strings.ReplaceAll(dto.appeal.Fname, "`", "")
	lname := strings.ReplaceAll(dto.appeal.Lname, "`", "")
	mname := strings.ReplaceAll(dto.appeal.Mname, "`", "")
	subject := strings.ReplaceAll(dto.appeal.Subject, "`", "")

	text += fmt.Sprintf(
		"*%s*: `%s %s %s`\n"+
			"*%s*: `%s`",
		s.translate(newCommandFullNameTranslateKey, dto.lang), lname, fname, mname,
		s.translate(newCommandSubjectTranslateKey, dto.lang), subject,
	)

	parseMode := markdownParseMode

	if err := s.api.SendMessage(entity.SendMessageAPIDTO{
		Text:   text,
		ChatID: dto.chatID,
		Options: &entity.Options{
			ParseMode: &parseMode,
		},
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("sendFinalAppeal - s.api.SendMessage error: ", err.Error())
	}

	return nil
}

type sendConfirmDTO struct {
	lang   string
	chatID int64
}

func (s *Service) sendConfirm(dto sendConfirmDTO) error {
	if err := s.api.SendMessageWithInlineKeyboard(entity.SendMessageWithInlineKeyboardAPIDTO{
		Text:   s.translate(newCommandConfirmSendAppealTranslateKey, dto.lang),
		ChatID: dto.chatID,
		Buttons: [][]entity.InlineButton{
			{
				{
					Text:         s.translate(newCommandSendConfirmOkButtonTranslateKey, dto.lang),
					CallbackData: confirmCallbackOk,
					URL:          "",
				},
				{
					Text:         s.translate(newCommandSendConfirmCancelButtonTranslateKey, dto.lang),
					CallbackData: confirmCallbackCancel,
					URL:          "",
				},
			},
		},
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("sendConfirm - s.api.SendMessage error: ", err.Error())
	}

	return nil
}

const (
	confirmCallbackOk     = "ok"
	confirmCallbackCancel = "cancel"
)

func (s *Service) NewCommandConfirmationState(ctx context.Context, dto entity.NewCommandConfirmationStateServiceDTO) error {
	switch dto.Data {
	case confirmCallbackOk:
		isDraft := false
		if err := s.repo.UpdateDraftAppeal(ctx, entity.UpdateAppealRepoDTO{
			UpdateAppealBase: entity.UpdateAppealBase{
				IsDraft: &isDraft,
			},
			ChatID: dto.ChatID,
		}); err != nil {
			s.log.WithFields(logrus.Fields{
				chatIDLoggingKey: dto.ChatID,
				domainLoggingKey: domainLoggingValue,
			}).Error("NewCommandConfirmationState - s.repo.UpdateDraftAppeal error: ", err.Error())

			s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
				lang:   dto.Lang,
				chatID: dto.ChatID,
			})

			return err
		}

		if err := s.api.SendMessage(entity.SendMessageAPIDTO{
			Text:    s.translate(newCommandConfirmAppealCreatedTranslateKey, dto.Lang),
			ChatID:  dto.ChatID,
			Options: nil,
		}); err != nil {
			s.log.WithFields(logrus.Fields{
				chatIDLoggingKey: dto.ChatID,
				domainLoggingKey: domainLoggingValue,
			}).Error("NewCommandConfirmationState - s.api.SendMessage error: ", err.Error())
		}

		return nil
	case confirmCallbackCancel:
		if err := s.repo.DeleteDraftAppeal(ctx, entity.DeleteDraftAppealRepoDTO{
			ChatID: dto.ChatID,
		}); err != nil {
			s.log.WithFields(logrus.Fields{
				chatIDLoggingKey: dto.ChatID,
				domainLoggingKey: domainLoggingValue,
			}).Error("NewCommandConfirmationState - s.repo.DeleteDraftAppeal error: ", err.Error())

			s.sendSomethingWentWrong(sendSomethingWentWrongDTO{
				lang:   dto.Lang,
				chatID: dto.ChatID,
			})

			return err
		}

		if err := s.api.SendMessage(entity.SendMessageAPIDTO{
			Text:    s.translate(newCommandConfirmAppealCanceledTranslateKey, dto.Lang),
			ChatID:  dto.ChatID,
			Options: nil,
		}); err != nil {
			s.log.WithFields(logrus.Fields{
				chatIDLoggingKey: dto.ChatID,
				domainLoggingKey: domainLoggingValue,
			}).Error("NewCommandConfirmationState - s.api.SendMessage error: ", err.Error())
		}

		return nil
	default:
		if err := s.api.SendMessage(entity.SendMessageAPIDTO{
			Text:    s.translate(newCommandConfirmAnswerNotExistsTranslateKey, dto.Lang),
			ChatID:  dto.ChatID,
			Options: nil,
		}); err != nil {
			s.log.WithFields(logrus.Fields{
				chatIDLoggingKey: dto.ChatID,
				domainLoggingKey: domainLoggingValue,
			}).Error("NewCommandConfirmationState - s.api.SendMessage error: ", err.Error())
		}

		return errs.ErrUnknownAnswer
	}
}
