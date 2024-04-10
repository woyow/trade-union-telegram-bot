package entity

type CreateChatCurrentStateServiceDTO struct {
	State  string
	ChatID int64
}

type CreateChatCurrentStateRepoDTO struct {
	State  string
	ChatID int64
}

type CreateChatCurrentStateOut struct {
	ID string
}

type SetChatCurrentStateServiceDTO struct {
	State  string
	ChatID int64
}

type SetChatCurrentStateRepoDTO struct {
	State  string
	ChatID int64
}

type GetChatCurrentStateServiceDTO struct {
	ChatID int64
}

type GetChatCurrentStateRepoDTO struct {
	ChatID int64
}

type GetChatCurrentStateOut struct {
	State string `bson:"state"`
}
