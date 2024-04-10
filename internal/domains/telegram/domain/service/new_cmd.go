package service

import (
	"context"
	"fmt"

	"github.com/NicoNex/echotron/v3"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"
)

func (s *Service) NewCommand(ctx context.Context, dto entity.NewCommandServiceDTO) error {
	if err := s.deleteDraftAppeal(ctx, deleteDraftAppealDTO{
		chatID: dto.ChatID,
	}); err != nil {
		if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommand - s.api.SendMessage error", err.Error())
		}
		return err
	}

	out, err := s.createAppeal(ctx, createAppealDTO{
		chatID:  dto.ChatID,
		isDraft: true,
	})
	if err != nil {
		if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommand - s.api.SendMessage error", err.Error())
		}
		return err
	} else {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Info("service: NewCommand - Draft appeal has been created with id: ", out.ID)
	}

	if _, err := s.api.SendMessage(s.translate(newCommandTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommand - s.api.SendMessage error", err.Error())
	}

	if _, err := s.api.SendMessage(s.translate(newCommandFirstNameTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommand - s.api.SendMessage error", err.Error())
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
		s.log.WithField(chatIDLoggingKey, dto.chatID).Error("service: deleteDraftAppeal - s.repo.DeleteDraftAppeal error: ", err.Error())
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
		s.log.WithField(chatIDLoggingKey, dto.chatID).Error("service: createAppeal - s.repo.CreateAppeal error: ", err.Error())
		return nil, err
	}

	return out, nil
}

func (s *Service) NewCommandFirstNameState(ctx context.Context, dto entity.NewCommandFirstNameStateServiceDTO) error {
	if err := s.repo.UpdateDraftAppeal(ctx, entity.UpdateAppealRepoDTO{
		UpdateAppealBase: entity.UpdateAppealBase{
			Fname: &dto.Text,
		},
		ChatID: dto.ChatID,
	}); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandFirstNameState - s.repo.UpdateDraftAppeal error: ", err.Error())

		if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandFirstNameState - s.api.SendMessage error", err.Error())
		}

		return err
	}

	if _, err := s.api.SendMessage(s.translate(newCommandLastNameTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandFirstNameState - s.api.SendMessage error", err.Error())
	}
	return nil
}

func (s *Service) NewCommandLastNameState(ctx context.Context, dto entity.NewCommandLastNameStateServiceDTO) error {
	if err := s.repo.UpdateDraftAppeal(ctx, entity.UpdateAppealRepoDTO{
		UpdateAppealBase: entity.UpdateAppealBase{
			Lname: &dto.Text,
		},
		ChatID: dto.ChatID,
	}); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandLastNameState - s.repo.UpdateDraftAppeal error: ", err.Error())

		if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandLastNameState - s.api.SendMessage error", err.Error())
		}

		return err
	}

	if _, err := s.api.SendMessage(s.translate(newCommandMiddleNameTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandLastNameState - s.api.SendMessage error", err.Error())
	}
	return nil
}

func (s *Service) NewCommandMiddleNameState(ctx context.Context, dto entity.NewCommandMiddleNameStateServiceDTO) error {
	if err := s.repo.UpdateDraftAppeal(ctx, entity.UpdateAppealRepoDTO{
		UpdateAppealBase: entity.UpdateAppealBase{
			Mname: &dto.Text,
		},
		ChatID: dto.ChatID,
	}); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandLastNameState - s.repo.UpdateDraftAppeal error: ", err.Error())

		if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandLastNameState - s.api.SendMessage error", err.Error())
		}

		return err
	}

	out, err := s.repo.GetDraftAppeal(ctx, entity.GetDraftAppealRepoDTO{
		ChatID: dto.ChatID,
	})
	if err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandLastNameState - s.repo.GetDraftAppeal error: ", err.Error())

		if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandLastNameState - s.api.SendMessage error", err.Error())
		}

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
	text += fmt.Sprintf(
		"*%s*: `%s %s %s`\n"+
			"*%s*: `Неизвестная тема обращения`",
		s.translate(newCommandFullNameTranslateKey, dto.lang), dto.appeal.Lname, dto.appeal.Fname, dto.appeal.Mname,
		s.translate(newCommandSubjectTranslateKey, dto.lang),
	)

	if _, err := s.api.SendMessage(text, dto.chatID, &echotron.MessageOptions{
		ParseMode: markdownParseMode,
	}); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.chatID).Error("service: NewCommandLastNameState - s.api.SendMessage error", err.Error())
	}

	return nil
}

type sendConfirmDTO struct {
	lang   string
	chatID int64
}

func (s *Service) sendConfirm(dto sendConfirmDTO) error {
	if _, err := s.api.SendMessage(s.translate(newCommandConfirmSendAppealTranslateKey, dto.lang), dto.chatID, &echotron.MessageOptions{
		ReplyMarkup: echotron.InlineKeyboardMarkup{
			InlineKeyboard: [][]echotron.InlineKeyboardButton{
				{
					{
						Text:         s.translate(newCommandSendConfirmOkButtonTranslateKey, dto.lang),
						CallbackData: confirmCallbackOk,
					},
					{
						Text:         s.translate(newCommandSendConfirmCancelButtonTranslateKey, dto.lang),
						CallbackData: confirmCallbackCancel,
					},
				},
			},
		},
	}); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.chatID).Error("service: sendConfirm - s.api.SendMessage error", err.Error())
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
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandConfirmationState - s.repo.UpdateDraftAppeal error: ", err.Error())

			if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
				s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandConfirmationState - s.api.SendMessage error", err.Error())
			}

			return err
		}

		if _, err := s.api.SendMessage(s.translate(newCommandConfirmAppealCreatedTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandConfirmationState - s.api.SendMessage error", err.Error())
		}

		return nil
	case confirmCallbackCancel:
		if err := s.repo.DeleteDraftAppeal(ctx, entity.DeleteDraftAppealRepoDTO{
			ChatID: dto.ChatID,
		}); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandConfirmationState - s.repo.DeleteDraftAppeal error: ", err.Error())

			if _, err := s.api.SendMessage(s.translate(somethingWentWrongTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
				s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandConfirmationState - s.api.SendMessage error", err.Error())
			}

			return err
		}

		if _, err := s.api.SendMessage(s.translate(newCommandConfirmAppealCanceledTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandConfirmationState - s.api.SendMessage error", err.Error())
		}

		return nil
	default:
		if _, err := s.api.SendMessage(s.translate(newCommandConfirmAnswerNotExistsTranslateKey, dto.Lang), dto.ChatID, nil); err != nil {
			s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: NewCommandConfirmationState - s.api.SendMessage error", err.Error())
		}
		return errs.ErrUnknownAnswer
	}

}
