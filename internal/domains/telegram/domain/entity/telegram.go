package entity

type Options struct {
	ParseMode *string
}

type SendMessageApiDTO struct {
	Text    string
	ChatID  int64
	Options *Options
}

type InlineButton struct {
	Text         string
	CallbackData string
	URL          string
}

type SendMessageWithInlineKeyboardApiDTO struct {
	Text    string
	ChatID  int64
	Buttons [][]InlineButton
}
