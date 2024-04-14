package service

const (
	langEn = "en"
	langRu = "ru"
)

const (
	// Common translate keys.
	somethingWentWrongTranslateKey = "something_went_wrong.message"

	// Unknown command translate keys.
	unknownCommandTranslateKey = "unknown_cmd.message"

	// Start command translate keys.
	startCommandTranslateKey = "start_cmd.message"

	// New command translate keys.
	newCommandTranslateKey                        = "new_cmd.message"
	newCommandFirstNameTranslateKey               = "new_cmd.type_first_name.message"
	newCommandLastNameTranslateKey                = "new_cmd.type_last_name.message"
	newCommandMiddleNameTranslateKey              = "new_cmd.type_middle_name.message"
	newCommandChooseSubjectOfAppealTranslateKey   = "new_cmd.type_choose_appeal_subject.text"
	newCommandFullNameTranslateKey                = "new_cmd.full_name.text"
	newCommandSubjectTranslateKey                 = "new_cmd.subject_of_appeal.text"
	newCommandSendConfirmOkButtonTranslateKey     = "new_cmd.confirm_ok.button"
	newCommandSendConfirmCancelButtonTranslateKey = "new_cmd.confirm_cancel.button"
	newCommandConfirmSendAppealTranslateKey       = "new_cmd.confirm_send_appeal.message"
	newCommandConfirmAnswerNotExistsTranslateKey  = "new_cmd.confirm_answer_not_exists.message"
	newCommandConfirmAppealCreatedTranslateKey    = "new_cmd.confirm_appeal_created.message"
	newCommandConfirmAppealCanceledTranslateKey   = "new_cmd.confirm_appeal_canceled.message"
)

type translateMap map[string]map[string]string

func getTranslateMap() translateMap {
	translate := translateMap{
		somethingWentWrongTranslateKey: map[string]string{
			langEn: "Oops. Something went wrong. Try again or later.",
			langRu: "Уупс. Что-то пошло не так. Попробуйте заново, либо позже.",
		},
		unknownCommandTranslateKey: map[string]string{
			langEn: "Unknown command. Try again!",
			langRu: "Неизвестная команда. Попробуйте снова!",
		},
		startCommandTranslateKey: map[string]string{
			langEn: "Hello, this bot assistant will help you draft a trade union appeal.",
			langRu: "Приветствую, этот бот-помощник поможет составить обращение профсоюз.",
		},
		newCommandTranslateKey: map[string]string{
			langEn: "To create a new request, the bot needs some data. Please enter them.",
			langRu: "Для создания нового обращения боту необходимы некоторые данные. Пожалуйста, введите их.",
		},
		newCommandFirstNameTranslateKey: map[string]string{
			langEn: "Enter your name:",
			langRu: "Введите имя:",
		},
		newCommandLastNameTranslateKey: map[string]string{
			langEn: "Enter last name:",
			langRu: "Введите фамилию:",
		},
		newCommandMiddleNameTranslateKey: map[string]string{
			langEn: "Enter middle name:",
			langRu: "Введите отчество:",
		},
		newCommandChooseSubjectOfAppealTranslateKey: map[string]string{
			langEn: "Select the subject of your appeal:",
			langRu: "Выберите тему обращения:",
		},
		newCommandConfirmSendAppealTranslateKey: map[string]string{
			langEn: "Confirm or cancel sending the request",
			langRu: "Подтвердите или отмените отправку обращения",
		},
		newCommandFullNameTranslateKey: map[string]string{
			langEn: "Full name",
			langRu: "ФИО",
		},
		newCommandSubjectTranslateKey: map[string]string{
			langEn: "Subject of appeal",
			langRu: "Тема обращения",
		},
		newCommandSendConfirmOkButtonTranslateKey: map[string]string{
			langEn: "Confirm ✅",
			langRu: "Подтверждаю ✅",
		},
		newCommandSendConfirmCancelButtonTranslateKey: map[string]string{
			langEn: "Cancel ❌",
			langRu: "Отмена ❌",
		},
		newCommandConfirmAnswerNotExistsTranslateKey: map[string]string{
			langEn: "There is no such answer. Try again.",
			langRu: "Такого ответа не существует. Попробуйте заново.",
		},
		newCommandConfirmAppealCreatedTranslateKey: map[string]string{
			langEn: "Appeal has been created.",
			langRu: "Обращение было создано.",
		},
		newCommandConfirmAppealCanceledTranslateKey: map[string]string{
			langEn: "Appeal has been canceled.",
			langRu: "Обращение было отменено.",
		},
	}
	return translate
}

const (
	translateNotExists = "translate not exists"
	defaultLang        = langRu
)

func (s *Service) translate(key, lang string) string {
	// Check language support
	switch lang {
	case langEn, langRu:
	default:
		lang = defaultLang
	}

	val, ok := s.translateDict[key][lang]
	if !ok {
		return translateNotExists
	}

	return val
}
