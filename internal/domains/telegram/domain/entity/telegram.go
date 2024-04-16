package entity

type Options struct {
	ParseMode *string
}

type SendMessageAPIDTO struct {
	Text    string
	ChatID  int64
	Options *Options
}

type InlineButton struct {
	Text         string
	CallbackData string
	URL          string
}

type SendMessageWithInlineKeyboardAPIDTO struct {
	Text    string
	ChatID  int64
	Buttons [][]InlineButton
}
