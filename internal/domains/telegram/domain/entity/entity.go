package entity

type HandleCommand struct {
	Lang   string
	ChatID int64
}

type HandleMessage struct {
	Lang   string
	Text   string
	ChatID int64
}

type HandleCallback struct {
	Lang   string
	Data   string
	ChatID int64
}

type UnknownCommandServiceDTO struct {
	HandleCommand
}
